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
		// No results on panic
		interpreter.ReturnResults = make([]symbolic.SymbolicExpression, 0)

		return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}
	case *ssa.Select:
		panic("Support me!")
	case *ssa.Store:
		val := interpreter.resolveExpression(i.Val)

		switch a := i.Addr.(type) {
		case *ssa.FieldAddr:
			base := interpreter.resolveExpression(a.X)
			br, ok := base.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Store: FieldAddr base is %T, want *symbolic.Ref", base))
			}
			interpreter.Heap.AssignField(br, a.Field, val)

		case *ssa.IndexAddr:
			// *(&arr[i]) = val  => arr[i] = val
			base := interpreter.resolveExpression(a.X)
			br, ok := base.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Store: IndexAddr base is %T, want *symbolic.Ref", base))
			}

			idxExpr := interpreter.resolveExpression(a.Index)
			ic, ok := idxExpr.(*symbolic.IntConstant)
			if !ok {
				panic(fmt.Sprintf("Store: Index must be const int, got %T (%s)", idxExpr, idxExpr.String()))
			}

			interpreter.Heap.AssignToArray(br, int(ic.Value), val)

		default:
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
			thenState := *interpreter // копия
			thenState.PathCondition = symbolic.NewBinaryOperation(
				thenState.PathCondition,
				cond,
				symbolic.AND,
			)
			out = append(out, &Item{value: thenState, block: block.Succs[0], instrIdx: 0, pred: block})
		}
		if len(i.Block().Succs) >= 2 {
			elseState := *interpreter // копия
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

	// Instructions that also implement Value interface will resolve later
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
	// A Parameter represents an input parameter of a function.
	case *ssa.Parameter:
		// struct / *struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			return interpreter.Heap.ParamStructRef(v.Name(), typeName, fields)
		}

		// []T
		if sl, ok := v.Type().Underlying().(*types.Slice); ok {
			elemTy := goTypeToSym(sl.Elem())
			return interpreter.Heap.ParamArrayRef(v.Name(), elemTy)
		}

		// *T (указатель на примитив/строку)
		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			elemTy := goTypeToSym(p.Elem())
			return interpreter.Heap.ParamRef(v.Name(), elemTy)
		}

		// обычные примитивы
		ty := goTypeToSym(v.Type())
		ref := interpreter.Heap.ParamRef(v.Name(), ty)
		return interpreter.Heap.GetValue(ref)

	/*
		The Alloc instruction:
		1. reserves space for a variable of the given type,
		2. zero-initializes it, TODO??? (In my case its z3.Value and solver decide which value it is when no assertations)
		3. yields its address.

		Alloc values are always addresses, and have pointer types, so the
		type of the allocated variable is actually
		Type().Underlying().(*types.Pointer).Elem().
	*/
	case *ssa.Alloc:
		var res symbolic.SymbolicExpression

		// struct / *struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			res = interpreter.Heap.AllocateStruct(typeName, fields)
			interpreter.Env[v] = res
			return res
		}

		// []T
		if sl, ok := v.Type().Underlying().(*types.Slice); ok {
			elemTy := goTypeToSym(sl.Elem())
			res = interpreter.Heap.AllocateArray(elemTy)
			interpreter.Env[v] = res
			return res
		}

		// pointer/array-like (твой текущий хак)
		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			elemTy := goTypeToSym(p.Elem())
			res = interpreter.Heap.AllocateArray(elemTy)
			interpreter.Env[v] = res
			return res
		}

		// Primitive
		ty := goTypeToSym(v.Type())
		ref := interpreter.Heap.Allocate(ty)
		res = interpreter.Heap.GetValue(ref)

		interpreter.Env[v] = res
		return res

	/*
		The Slice instruction yields a slice of an existing string, slice
		or *array X between optional integer bounds Low and High.

		Type() returns string if the type of X was string, otherwise a
		*types.Slice with the same element type as X.
	*/
	case *ssa.Slice:
		res := interpreter.resolveExpression(v.X) // TODO Support slices
		interpreter.Env[v] = res
		return res

	case *ssa.MakeSlice:
		sl, ok := v.Type().Underlying().(*types.Slice)
		if !ok {
			panic(fmt.Sprintf("MakeSlice: unexpected type %T (%s)", v.Type().Underlying(), v.Type().String()))
		}

		elemTy := goTypeToSym(sl.Elem())
		res := interpreter.Heap.AllocateArray(elemTy)

		interpreter.Env[v] = res
		return res

	// IntConstant, BoolConstant, StringConstant
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
			panic(fmt.Sprintf("unsupported value %T", v.Value.Kind()))
		}

	// X Op Y
	/*
		One of:
		ADD SUB MUL QUO REM          + - * / %
		AND OR XOR SHL SHR AND_NOT   & | ^ << >> &^
		EQL NEQ LSS LEQ GTR GEQ      == != < <= < >=
	*/
	case *ssa.BinOp:
		l := interpreter.resolveExpression(v.X)
		r := interpreter.resolveExpression(v.Y)

		if sop, ok := tokenToBinOp(v.Op); ok {
			return symbolic.NewBinaryOperation(l, r, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	// Op X
	/*
		One of:
		NOT SUB ARROW MUL XOR ! - <- * ^
	*/
	case *ssa.UnOp:
		// 1) СНАЧАЛА special-case для load: *X
		if v.Op == token.MUL {
			// &foo2.a
			if fa, ok := v.X.(*ssa.FieldAddr); ok {
				base := interpreter.resolveExpression(fa.X) // fa.X = foo2
				br, ok := base.(*symbolic.Ref)
				if !ok {
					panic(fmt.Sprintf("Load: FieldAddr base is %T, want *symbolic.Ref", base))
				}
				return interpreter.Heap.GetFieldValue(br, fa.Field)
			}

			// &foo2[Index]
			if ia, ok := v.X.(*ssa.IndexAddr); ok {
				base := interpreter.resolveExpression(ia.X)
				br, ok := base.(*symbolic.Ref)
				if !ok {
					panic(fmt.Sprintf("Load: IndexAddr base is %T, want *symbolic.Ref", base))
				}

				idxExpr := interpreter.resolveExpression(ia.Index)
				ic, ok := idxExpr.(*symbolic.IntConstant)
				if !ok {
					panic(fmt.Sprintf("Load: Index must be const int, got %T (%s)", idxExpr, idxExpr.String()))
				}

				return interpreter.Heap.GetFromArray(br, int(ic.Value))
			}

			// Primitive ref
			addrExpr := interpreter.resolveExpression(v.X)
			r, ok := addrExpr.(*symbolic.Ref)
			if !ok {
				panic(fmt.Sprintf("Load: %T is not *symbolic.Ref", addrExpr))
			}
			return interpreter.Heap.GetValue(r)
		}

		// 2) ТОЛЬКО ПОТОМ резолвим операнд для остальных унарных операций
		operand := interpreter.resolveExpression(v.X)

		if sop, ok := tokenToUnOp(v.Op); ok {
			return symbolic.NewUnaryOperation(operand, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	// The Call instruction represents a function or method call.
	case *ssa.Call:
		panic("Implement me HW5!")

	default:
		// fallback: просто символ по имени SSA value
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
		panic("Not impl!")
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
		panic(fmt.Sprintf("Not impl! %s", op.String()))
	}
}

// TODO Внести всю логику выведения типов в эту функцию и не размазывать по интерпретатору
func goTypeToSym(t types.Type) symbolic.ExpressionType {
	if t == nil {
		panic("nil")
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
		panic(fmt.Sprintf("Not impl: %T\n", tt))
	}

	panic(fmt.Sprintf("Not impl: %T\n", t))
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
