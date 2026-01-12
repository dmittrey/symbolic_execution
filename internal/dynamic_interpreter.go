package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"

	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"

	"golang.org/x/tools/go/ssa"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	ReturnResults []symbolic.SymbolicExpression
	Heap          memory.Memory
	Env           map[ssa.Value]symbolic.SymbolicExpression
}

func NewInterpreter(analyser *Analyser, pathCond symbolic.SymbolicExpression) Interpreter {
	return Interpreter{
		CallStack:     make([]CallStackFrame, 0),
		Analyser:      analyser,
		PathCondition: pathCond,
		Heap:          memory.NewSymbolicMemory(),
		Env:           make(map[ssa.Value]symbolic.SymbolicExpression),
	}
}

type CallStackFrame struct {
	Function    *ssa.Function
	LocalMemory map[string]symbolic.SymbolicExpression
	ReturnValue symbolic.SymbolicExpression
}

func (interpreter *Interpreter) interpretDynamically(block *ssa.BasicBlock, instrIdx int, pred *ssa.BasicBlock) []*Item {
	if block == nil || instrIdx < 0 || instrIdx >= len(block.Instrs) {
		return nil
	}

	instr := block.Instrs[instrIdx]
	fmt.Printf("interpretDynamically: %s\n", instr.String())

	if instr == nil {
		fmt.Println("interpretDynamically with nil!")
		return nil
	}

	switch i := instr.(type) {

	case *ssa.Return:
		interpreter.ReturnResults = make([]symbolic.SymbolicExpression, len(i.Results))
		for idx, val := range i.Results {
			interpreter.ReturnResults[idx] = interpreter.resolveExpression(val)
		}
		return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}

	case *ssa.Panic:
		interpreter.ReturnResults = make([]symbolic.SymbolicExpression, 0)
		return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}

	case *ssa.Select:
		panic("Support me!")

	case *ssa.Store:
		val := interpreter.resolveExpression(i.Val)

		switch a := i.Addr.(type) {

		case *ssa.FieldAddr:
			// *(&obj.field) = val  => obj.field = val
			base := interpreter.resolveExpression(a.X)
			br, ok := base.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Store: FieldAddr base is %T, want *symbolic.Ref", base))
			}
			fr := interpreter.Heap.FieldAddress(br, a.Field)
			interpreter.Heap.AssignValue(fr, val)

		case *ssa.IndexAddr:
			// *(&arr[i]) = val  => arr[i] = val
			base := interpreter.resolveExpression(a.X)
			br, ok := base.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Store: IndexAddr base is %T, want *symbolic.Ref", base))
			}

			idxExpr := interpreter.resolveExpression(a.Index)
			if idxExpr.Type() != symbolic.IntType {
				panic(fmt.Sprintf("Store: Index must be int, got %T (%s)", idxExpr, idxExpr.String()))
			}
			interpreter.Heap.AssignToArray(br, idxExpr, val)

		default:
			// обычное присваивание по адресу ref
			addrExpr := interpreter.resolveExpression(i.Addr)
			r, ok := addrExpr.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Store: Addr is %T, want *symbolic.Ref", addrExpr))
			}
			interpreter.Heap.AssignValue(r, val)
		}

		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	case *ssa.Jump:
		out := make([]*Item, 0, len(block.Succs))
		for _, s := range block.Succs {
			out = append(out, &Item{value: *interpreter, block: s, instrIdx: 0, pred: block})
		}
		return out

	case *ssa.If:
		cond := interpreter.resolveExpression(i.Cond)

		out := make([]*Item, 0, 2)

		// then -> Succs[0]
		if len(i.Block().Succs) >= 1 {
			thenState := *interpreter
			thenState.PathCondition = symbolic.NewBinaryOperation(
				thenState.PathCondition,
				cond,
				symbolic.AND,
			)
			out = append(out, &Item{value: thenState, block: block.Succs[0], instrIdx: 0, pred: block})
		}

		// else -> Succs[1]
		if len(i.Block().Succs) >= 2 {
			elseState := *interpreter
			notCond := symbolic.NewUnaryOperation(cond, symbolic.NOT)
			elseState.PathCondition = symbolic.NewBinaryOperation(
				elseState.PathCondition,
				notCond,
				symbolic.AND,
			)
			out = append(out, &Item{value: elseState, block: block.Succs[1], instrIdx: 0, pred: block})
		}
		return out

	case *ssa.Phi:
		edge, ok := pickPhiEdge(i, block, pred)
		if !ok {
			panic("phi: can't pick edge (pred is nil or not found)")
		}
		phiVal := interpreter.resolveExpression(edge)
		interpreter.Env[i] = phiVal
		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	// Инструкции, которые также Value — будут резолвиться в resolveExpression
	case *ssa.BinOp, *ssa.UnOp,
		*ssa.FieldAddr, *ssa.IndexAddr,
		*ssa.Slice, *ssa.MakeSlice, *ssa.Alloc:
		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	default:
		panic(fmt.Sprintf("Support %T!", i))
	}
}

func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {
	if value == nil {
		fmt.Println("resolveExpression with nil!")
		return nil
	}

	fmt.Printf("resolveExpression: %s\n", value.String())

	if e, ok := interpreter.Env[value]; ok {
		return e
	}

	switch v := value.(type) {

	// Parameter: входной параметр функции
	case *ssa.Parameter:
		// struct / *struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			// моделируем как "адрес на struct"
			return interpreter.Heap.ParamStruct(typeName, fields)
		}

		// []T / [][]T / ...
		if depth, base, ok := sliceDepthBase(v.Type()); ok {
			// []struct — разрешаем только одномерный
			if typeName, fields, ok2 := structMeta(base); ok2 {
				if depth != 1 {
					panic("Parameter: multidimensional array of structs is not supported")
				}
				return interpreter.Heap.ParamStructsArray(0, typeName, fields)
			}

			baseTy := goTypeToSym(base)
			dim := depth - 1 // []T => 0, [][]T => 1, ...
			if dim == 0 {
				return interpreter.Heap.ParamPrimitiveArray(baseTy)
			}
			return interpreter.Heap.ParamComplexArray(dim, baseTy)
		}

		// *T (указатель на примитив/строку)
		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			// *struct уже поймали выше в structMeta
			elemTy := goTypeToSym(p.Elem())
			return interpreter.Heap.Param(elemTy) // адрес
		}

		// обычный примитив как value: создаём адрес + читаем значение
		ty := goTypeToSym(v.Type())
		ref := interpreter.Heap.Param(ty)
		return interpreter.Heap.GetValue(ref)

	/*
		Alloc:
		- всегда возвращает адрес (Ref).
	*/
	case *ssa.Alloc:
		// struct / *struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			r := interpreter.Heap.AllocateStruct(typeName, fields)
			interpreter.Env[v] = r
			return r
		}

		// []T / [][]T ...
		if depth, base, ok := sliceDepthBase(v.Type()); ok {
			if typeName, fields, ok2 := structMeta(base); ok2 {
				if depth != 1 {
					panic("Alloc: multidimensional array of structs is not supported")
				}
				r := interpreter.Heap.AllocateStructsArray(0, typeName, fields)
				interpreter.Env[v] = r
				return r
			}

			baseTy := goTypeToSym(base)
			dim := depth - 1
			if dim == 0 {
				r := interpreter.Heap.AllocatePrimitiveArray(baseTy)
				interpreter.Env[v] = r
				return r
			}
			r := interpreter.Heap.AllocateComplexArray(dim, baseTy)
			interpreter.Env[v] = r
			return r
		}

		// pointer alloc: тип = *T
		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			if typeName, fields, ok2 := structMeta(p.Elem()); ok2 {
				r := interpreter.Heap.AllocateStruct(typeName, fields)
				interpreter.Env[v] = r
				return r
			}
			ty := goTypeToSym(p.Elem())
			r := interpreter.Heap.Allocate(ty)
			interpreter.Env[v] = r
			return r
		}

		// примитив
		ty := goTypeToSym(v.Type())
		r := interpreter.Heap.Allocate(ty)
		interpreter.Env[v] = r
		return r

	/*
		Slice(X[low:high]) — пока хак: просто возвращаем X
	*/
	case *ssa.Slice:
		res := interpreter.resolveExpression(v.X)
		interpreter.Env[v] = res
		return res

	/*
		MakeSlice: создаём новый массив/слайс
	*/
	case *ssa.MakeSlice:
		depth, base, ok := sliceDepthBase(v.Type())
		if !ok {
			panic(fmt.Sprintf("MakeSlice: unexpected type %T (%s)", v.Type().Underlying(), v.Type().String()))
		}

		if typeName, fields, ok2 := structMeta(base); ok2 {
			if depth != 1 {
				panic("MakeSlice: multidimensional array of structs is not supported")
			}
			r := interpreter.Heap.AllocateStructsArray(0, typeName, fields)
			interpreter.Env[v] = r
			return r
		}

		baseTy := goTypeToSym(base)
		dim := depth - 1
		if dim == 0 {
			r := interpreter.Heap.AllocatePrimitiveArray(baseTy)
			interpreter.Env[v] = r
			return r
		}
		r := interpreter.Heap.AllocateComplexArray(dim, baseTy)
		interpreter.Env[v] = r
		return r

	// Константы
	case *ssa.Const:
		if v.Value == nil {
			panic("unsupported nil value")
		}

		switch v.Value.Kind() {
		case constant.Bool:
			return symbolic.NewBoolConstant(constant.BoolVal(v.Value))

		case constant.Int:
			i, ok := constant.Int64Val(v.Value)
			if !ok {
				panic(fmt.Sprintf("big int const not supported: %s", v.Value.String()))
			}
			return symbolic.NewIntConstant(i)

		case constant.String:
			return symbolic.NewStringConstant(constant.StringVal(v.Value))

		default:
			panic(fmt.Sprintf("unsupported const kind: %v", v.Value.Kind()))
		}

	// X Op Y
	case *ssa.BinOp:
		l := interpreter.resolveExpression(v.X)
		r := interpreter.resolveExpression(v.Y)

		if sop, ok := tokenToBinOp(v.Op); ok {
			return symbolic.NewBinaryOperation(l, r, sop)
		}

		// fallback
		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	// Op X
	case *ssa.UnOp:
		// special-case load: *X
		if v.Op == token.MUL {
			// *(&obj.field)
			if fa, ok := v.X.(*ssa.FieldAddr); ok {
				base := interpreter.resolveExpression(fa.X)
				br, ok := base.(*symbolic.Ref)
				if !ok {
					panic(fmt.Sprintf("Load: FieldAddr base is %T, want *symbolic.Ref", base))
				}
				fr := interpreter.Heap.FieldAddress(br, fa.Field)
				return interpreter.Heap.GetValue(fr)
			}

			// *(&arr[i])
			if ia, ok := v.X.(*ssa.IndexAddr); ok {
				base := interpreter.resolveExpression(ia.X)
				br, ok := base.(*symbolic.Ref)
				if !ok {
					panic(fmt.Sprintf("Load: IndexAddr base is %T, want *symbolic.Ref", base))
				}

				idxExpr := interpreter.resolveExpression(ia.Index)
				if idxExpr.Type() != symbolic.IntType {
					panic(fmt.Sprintf("Load: Index must be int, got %T (%s)", idxExpr, idxExpr.String()))
				}
				fmt.Printf("IndexAddr: base=%s, baseRefType=%s, dim=%d, baseType=%s\n",
					br.String(), br.RefType.String(), br.ArrayDimension, br.ArrayBaseType.String())

				return interpreter.Heap.GetFromArray(br, idxExpr)
			}

			// *ref (примитив или field-ref)
			addrExpr := interpreter.resolveExpression(v.X)
			r, ok := addrExpr.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Load: %T is not *symbolic.Ref", addrExpr))
			}
			return interpreter.Heap.GetValue(r)
		}

		operand := interpreter.resolveExpression(v.X)

		if sop, ok := tokenToUnOp(v.Op); ok {
			return symbolic.NewUnaryOperation(operand, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	case *ssa.Call:
		panic("Implement me HW5!")

	default:
		panic(fmt.Sprintf("Support me! %T", v))
	}
}

func tokenToBinOp(op token.Token) (symbolic.BinaryOperator, bool) {
	switch op {
	case token.ADD:
		return symbolic.ADD, true
	case token.SUB:
		return symbolic.SUB, true
	case token.MUL:
		return symbolic.MUL, true
	case token.QUO:
		return symbolic.DIV, true
	case token.REM:
		return symbolic.REM, true

	case token.AND:
		return symbolic.AND, true
	case token.OR:
		return symbolic.OR, true
	case token.XOR:
		return symbolic.XOR, true
	case token.SHL:
		return symbolic.SHL, true
	case token.SHR:
		return symbolic.SHR, true
	case token.AND_NOT:
		return symbolic.AND_NOT, true

	case token.EQL:
		return symbolic.EQ, true
	case token.NEQ:
		return symbolic.NE, true
	case token.LSS:
		return symbolic.LT, true
	case token.LEQ:
		return symbolic.LE, true
	case token.GTR:
		return symbolic.GT, true
	case token.GEQ:
		return symbolic.GE, true
	default:
		return 0, false
	}
}

func tokenToUnOp(op token.Token) (symbolic.UnaryOperator, bool) {
	switch op {
	case token.NOT:
		return symbolic.NOT, true
	case token.SUB:
		return symbolic.NEG, true
	case token.XOR:
		return symbolic.UXOR, true
	default:
		return 0, false
	}
}

func goTypeToSym(t types.Type) symbolic.ExpressionType {
	if t == nil {
		panic("nil type")
	}

	switch tt := t.Underlying().(type) {
	case *types.Basic:
		switch {
		case tt.Info()&types.IsBoolean != 0:
			return symbolic.BoolType
		case tt.Info()&types.IsInteger != 0:
			return symbolic.IntType
		case tt.Info()&types.IsString != 0:
			return symbolic.StringType
		}
	case *types.Pointer:
		return goTypeToSym(tt.Elem())
	case *types.Struct:
		return symbolic.StructType
	case *types.Array:
		return goTypeToSym(tt.Elem())
	default:
		panic(fmt.Sprintf("goTypeToSym: not impl: %T (%s)", tt, t.String()))
	}

	panic(fmt.Sprintf("goTypeToSym: not impl: %T (%s)", t, t.String()))
}

func nextInSameBlock(st Interpreter, block *ssa.BasicBlock, instrIdx int, pred *ssa.BasicBlock) []*Item {
	nxt := instrIdx + 1
	if nxt < len(block.Instrs) {
		return []*Item{{value: st, block: block, instrIdx: nxt, pred: pred}}
	}
	out := make([]*Item, 0, len(block.Succs))
	for _, s := range block.Succs {
		out = append(out, &Item{value: st, block: s, instrIdx: 0, pred: block})
	}
	return out
}

func pickPhiEdge(phi *ssa.Phi, block *ssa.BasicBlock, pred *ssa.BasicBlock) (ssa.Value, bool) {
	if pred == nil {
		return nil, false
	}
	for idx, p := range block.Preds {
		if p == pred {
			if idx < 0 || idx >= len(phi.Edges) {
				return nil, false
			}
			return phi.Edges[idx], true
		}
	}
	return nil, false
}

// sliceDepthBase: извлекаем глубину слайса и базовый тип.
//
// []T      => depth=1, base=T
// [][]T    => depth=2, base=T
// [][][]T  => depth=3, base=T
//
// Указатели сверху снимаем: *[][]T тоже ок.
func sliceDepthBase(t types.Type) (depth int, base types.Type, ok bool) {
	for {
		if p, isPtr := t.(*types.Pointer); isPtr {
			t = p.Elem()
			continue
		}
		break
	}

	for {
		sl, isSlice := t.Underlying().(*types.Slice)
		if !isSlice {
			break
		}
		depth++
		t = sl.Elem()
	}

	if depth == 0 {
		return 0, nil, false
	}
	return depth, t, true
}

// structMeta: возвращает имя структуры и типы полей (символьные).
// Разворачивает указатели и named-типы.
func structMeta(t types.Type) (typeName string, fieldTypes []symbolic.ExpressionType, ok bool) {
	for {
		if p, isPtr := t.(*types.Pointer); isPtr {
			t = p.Elem()
			continue
		}
		break
	}

	if n, isNamed := t.(*types.Named); isNamed {
		typeName = n.Obj().Name()
		t = n.Underlying()
	}

	st, isStruct := t.(*types.Struct)
	if !isStruct {
		return "", nil, false
	}

	fieldTypes = make([]symbolic.ExpressionType, 0, st.NumFields())
	for i := 0; i < st.NumFields(); i++ {
		ft := st.Field(i).Type()
		fieldTypes = append(fieldTypes, goTypeToSym(ft))
	}

	if typeName == "" {
		typeName = "anon_struct"
	}
	return typeName, fieldTypes, true
}
