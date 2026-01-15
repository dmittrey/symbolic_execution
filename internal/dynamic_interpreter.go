package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"unsafe"

	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"

	"golang.org/x/tools/go/ssa"
)

type LValueKind int

const (
	LVDirectRef   LValueKind = iota // обычный *Ref (скаляр/struct)
	LVArrayElem                     // x[i]
	LVStructField                   // obj.field
)

type LValue struct {
	Kind  LValueKind
	Base  *symbolic.Ref               // для DirectRef: сам ref; для ArrayElem/StructField: base-object
	Index symbolic.SymbolicExpression // только для ArrayElem
	Field int                         // только для StructField

	SsaType types.Type
}

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	ReturnResults []symbolic.SymbolicExpression
	Heap          memory.Memory
	Env           map[ssa.Value]symbolic.SymbolicExpression

	LenCache map[string]symbolic.SymbolicExpression
	paramRef map[ssa.Value]*symbolic.Ref
}

func NewInterpreter(analyser *Analyser, pathCond symbolic.SymbolicExpression) Interpreter {
	return Interpreter{
		CallStack:     make([]CallStackFrame, 0),
		Analyser:      analyser,
		PathCondition: pathCond,
		Heap:          memory.NewSymbolicMemory(),
		Env:           make(map[ssa.Value]symbolic.SymbolicExpression),

		LenCache: make(map[string]symbolic.SymbolicExpression),
		paramRef: make(map[ssa.Value]*symbolic.Ref),
	}
}

type CallStackFrame struct {
	Function    *ssa.Function
	LocalMemory map[string]symbolic.SymbolicExpression
	ReturnValue symbolic.SymbolicExpression

	// To save place we should return after nested function
	CallerBlock    *ssa.BasicBlock
	CallerInstrIdx int
	CallerPred     *ssa.BasicBlock
	CallInstr      ssa.Value
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
		// return to caller
		if len(interpreter.CallStack) > 0 {
			retVals := make([]symbolic.SymbolicExpression, len(i.Results))
			for idx, v := range i.Results {
				retVals[idx] = interpreter.resolveExpression(v)
			}

			var ret symbolic.SymbolicExpression
			if len(retVals) == 0 {
				ret = symbolic.NewNilConstant()
			} else {
				ret = retVals[0]
			}

			// pop frame
			top := interpreter.CallStack[len(interpreter.CallStack)-1]
			interpreter.CallStack = interpreter.CallStack[:len(interpreter.CallStack)-1]

			// bind return value to call instruction in caller env
			if top.CallInstr != nil {
				interpreter.Env[top.CallInstr] = ret
			}

			// continue in caller at next instruction after call
			if top.CallerBlock == nil {
				return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}
			}
			return []*Item{{value: *interpreter, block: top.CallerBlock, instrIdx: top.CallerInstrIdx, pred: top.CallerPred}}
		}

		// top-level return
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

		lv := interpreter.resolveAddr(i.Addr)
		interpreter.store(lv, val)

		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	case *ssa.Jump:
		out := make([]*Item, 0, len(block.Succs))
		if len(block.Succs) <= 1 {
			for _, s := range block.Succs {
				out = append(out, &Item{value: *interpreter, block: s, instrIdx: 0, pred: block})
			}
			return out
		}

		for _, s := range block.Succs {
			st := interpreter.cloneMapsOnly()
			out = append(out, &Item{value: st, block: s, instrIdx: 0, pred: block})
		}
		return out

	case *ssa.If:
		cond := interpreter.resolveExpression(i.Cond)

		out := make([]*Item, 0, 2)

		// then -> Succs[0]
		if len(i.Block().Succs) >= 1 {
			thenState := interpreter.cloneMapsOnly()
			thenState.PathCondition = symbolic.NewBinaryOperation(
				thenState.PathCondition,
				cond,
				symbolic.AND,
			)
			out = append(out, &Item{value: thenState, block: block.Succs[0], instrIdx: 0, pred: block})
		}

		// else -> Succs[1]
		if len(i.Block().Succs) >= 2 {
			elseState := interpreter.cloneMapsOnly()
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

	case *ssa.Call:
		// Take from resolveExpression
		if _, ok := i.Call.Value.(*ssa.Builtin); ok {
			interpreter.Env[i] = interpreter.resolveExpression(i)
			return nextInSameBlock(*interpreter, block, instrIdx, pred)
		}

		// interprocedural analysis
		if f, ok := i.Call.Value.(*ssa.Function); ok {
			// SymVar if external
			if f == nil || len(f.Blocks) == 0 {
				stub := symbolic.NewSymbolicVariable("call_"+sanitizeName(i.String()), goTypeToSym(i.Type()))
				interpreter.Env[i] = stub
				return nextInSameBlock(*interpreter, block, instrIdx, pred)
			}

			frame := CallStackFrame{
				Function:       f,
				LocalMemory:    make(map[string]symbolic.SymbolicExpression),
				CallerBlock:    block,
				CallerInstrIdx: instrIdx + 1,
				CallerPred:     pred,
				CallInstr:      i,
			}

			// bind args -> callee params by name
			for pi, p := range f.Params {
				if pi >= len(i.Call.Args) {
					break
				}

				argExpr := interpreter.resolveExpression(i.Call.Args[pi])

				frame.LocalMemory[p.Name()] = argExpr
			}

			interpreter.CallStack = append(interpreter.CallStack, frame)

			entry := f.Blocks[0]
			return []*Item{{value: *interpreter, block: entry, instrIdx: 0, pred: nil}}
		}

		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	// Инструкции, которые также Value — будут резолвиться в resolveExpression
	case *ssa.BinOp, *ssa.UnOp,
		*ssa.FieldAddr, *ssa.IndexAddr,
		*ssa.Slice, *ssa.MakeSlice, *ssa.Alloc,
		*ssa.Convert:
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

	case *ssa.Parameter:
		// Lookup in stack memory first
		if len(interpreter.CallStack) > 0 {
			for _, val := range interpreter.CallStack {
				if bound, ok := val.LocalMemory[v.Name()]; ok {
					return bound
				}
			}
		}

		if ref, ok := interpreter.paramRef[v]; ok {
			// pointer param => address
			if _, isPtr := v.Type().Underlying().(*types.Pointer); isPtr {
				return ref
			}
			// slice param => handle (array ref)
			if _, _, isSlice := sliceDepthBase(v.Type()); isSlice {
				return ref
			}
			// struct-by-value param => we model as address to struct
			if _, _, isStruct := structMeta(v.Type()); isStruct {
				return ref
			}
			// primitive-by-value param => read from cell
			return interpreter.Heap.GetValue(ref)
		}

		// create once

		// struct / *struct  => address to struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			r := interpreter.Heap.AllocateStruct(typeName, fields, true)
			interpreter.paramRef[v] = r
			return r
		}

		// []T / [][]T / ...
		if depth, base, ok := sliceDepthBase(v.Type()); ok {
			// []struct only 1D
			if typeName, fields, ok2 := structMeta(base); ok2 {
				if depth != 1 {
					panic("Parameter: multidimensional array of structs is not supported")
				}
				r := interpreter.Heap.AllocateStructsArray(0, typeName, fields, true)
				interpreter.paramRef[v] = r
				return r
			}

			baseTy := goTypeToSym(base)
			dim := depth - 1 // []T=>0, [][]T=>1, ...
			r := interpreter.Heap.AllocateArray(dim, baseTy, true)
			interpreter.paramRef[v] = r
			return r
		}

		// *T pointer to primitive/string/float/bool
		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			// *struct already caught above by structMeta
			elemTy := goTypeToSym(p.Elem())
			r := interpreter.Heap.Allocate(elemTy, true)
			interpreter.paramRef[v] = r
			return r
		}

		// primitive-by-value
		ty := goTypeToSym(v.Type())
		ref := interpreter.Heap.Allocate(ty, true)
		interpreter.paramRef[v] = ref
		return interpreter.Heap.GetValue(ref)

	/*
		Alloc:
		- всегда возвращает адрес (Ref).
	*/
	case *ssa.Alloc:
		// struct / *struct
		if typeName, fields, ok := structMeta(v.Type()); ok {
			r := interpreter.Heap.AllocateStruct(typeName, fields, false)
			interpreter.Env[v] = r
			return r
		}

		// []T / [][]T ...
		if depth, base, ok := sliceDepthBase(v.Type()); ok {
			if typeName, fields, ok2 := structMeta(base); ok2 {
				if depth != 1 {
					panic("Alloc: multidimensional array of structs is not supported")
				}
				r := interpreter.Heap.AllocateStructsArray(0, typeName, fields, false)
				interpreter.Env[v] = r
				return r
			}

			baseTy := goTypeToSym(base)
			dim := depth - 1
			r := interpreter.Heap.AllocateArray(dim, baseTy, false)
			interpreter.Env[v] = r
			return r
		}

		if p, ok := v.Type().Underlying().(*types.Pointer); ok {
			if typeName, fields, ok2 := structMeta(p.Elem()); ok2 {
				r := interpreter.Heap.AllocateStruct(typeName, fields, false)
				interpreter.Env[v] = r
				return r
			}
			ty := goTypeToSym(p.Elem())
			r := interpreter.Heap.Allocate(ty, false)
			interpreter.Env[v] = r
			return r
		}

		ty := goTypeToSym(v.Type())
		r := interpreter.Heap.Allocate(ty, false)
		interpreter.Env[v] = r
		return r

	case *ssa.Slice:
		res := interpreter.resolveExpression(v.X)
		interpreter.Env[v] = res
		return res

	case *ssa.Convert:
		src := interpreter.resolveExpression(v.X)
		dstTy := goTypeToSym(v.Type())

		if src.Type() == dstTy {
			interpreter.Env[v] = src
			return src
		}

		panic(fmt.Sprintf("Convert not supported: %s -> %s (%s)",
			src.Type().String(), dstTy.String(), v.String()))

	case *ssa.MakeSlice:
		depth, base, ok := sliceDepthBase(v.Type())
		if !ok {
			panic(fmt.Sprintf("MakeSlice: unexpected type %T (%s)", v.Type().Underlying(), v.Type().String()))
		}

		if typeName, fields, ok2 := structMeta(base); ok2 {
			if depth != 1 {
				panic("MakeSlice: multidimensional array of structs is not supported")
			}
			r := interpreter.Heap.AllocateStructsArray(0, typeName, fields, false)
			interpreter.Env[v] = r
			return r
		}

		baseTy := goTypeToSym(base)
		dim := depth - 1
		r := interpreter.Heap.AllocateArray(dim, baseTy, false)
		interpreter.Env[v] = r
		return r

	case *ssa.Const:
		if v.Value == nil {
			return interpreter.nilForType(v.Type())
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
		case constant.Float:
			f, ok := constant.Float64Val(v.Value)
			if !ok {
				panic(fmt.Sprintf("float const not supported: %s", v.Value.String()))
			}
			return symbolic.NewFloatConstant(f)
		default:
			panic(fmt.Sprintf("unsupported const kind: %v", v.Value.Kind()))
		}

	case *ssa.BinOp:
		l := interpreter.resolveExpression(v.X)
		r := interpreter.resolveExpression(v.Y)

		if sop, ok := tokenToBinOp(v.Op); ok {
			return symbolic.NewBinaryOperation(l, r, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	case *ssa.UnOp:
		if v.Op == token.MUL {
			lv := interpreter.resolveAddr(v.X)
			return interpreter.load(lv)
		}

		operand := interpreter.resolveExpression(v.X)

		if sop, ok := tokenToUnOp(v.Op); ok {
			return symbolic.NewUnaryOperation(operand, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	case *ssa.Call:
		if b, ok := v.Call.Value.(*ssa.Builtin); ok && b.Name() == "len" {
			if len(v.Call.Args) != 1 {
				panic("len(): unexpected args count")
			}

			arg := v.Call.Args[0]
			if c, ok := arg.(*ssa.Const); ok && c.Value == nil {
				return symbolic.NewIntConstant(0)
			}

			key := "len@" + arg.String()
			if cached, ok := interpreter.LenCache[key]; ok {
				interpreter.Env[v] = cached
				return cached
			}

			stub := symbolic.NewSymbolicVariable("len_"+sanitizeName(arg.String()), symbolic.IntType)
			interpreter.LenCache[key] = stub
			interpreter.Env[v] = stub
			return stub
		}

		// fallback
		stub := symbolic.NewSymbolicVariable("call_"+sanitizeName(v.String()), goTypeToSym(v.Type()))
		return stub

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
		case tt.Info()&types.IsFloat != 0:
			return symbolic.FloatType
		}

	case *types.Pointer:
		return goTypeToSym(tt.Elem())

	case *types.Struct:
		return symbolic.StructType

	case *types.Array:
		return goTypeToSym(tt.Elem())

	case *types.Slice:
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
	if len(block.Succs) <= 1 {
		for _, s := range block.Succs {
			out = append(out, &Item{value: st, block: s, instrIdx: 0, pred: block})
		}
		return out
	}

	for _, s := range block.Succs {
		cp := st
		cloned := (&cp).cloneMapsOnly()
		out = append(out, &Item{value: cloned, block: s, instrIdx: 0, pred: block})
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

func (it *Interpreter) resolveAddr(v ssa.Value) LValue {
	switch a := v.(type) {

	case *ssa.FieldAddr:
		base := it.resolveExpression(a.X)
		br, ok := base.(*symbolic.Ref)
		if !ok {
			panic("FieldAddr base not ref")
		}
		// важно: даже если br "не struct" по RefType — SSA скажет что это struct/*struct.
		return LValue{Kind: LVStructField, Base: br, Field: a.Field, SsaType: a.Type()}

	case *ssa.IndexAddr:
		base := it.resolveExpression(a.X)
		br, ok := base.(*symbolic.Ref)
		if !ok {
			panic("IndexAddr base not ref")
		}

		idx := it.resolveExpression(a.Index)
		if idx.Type() != symbolic.IntType {
			panic("Index must be int")
		}

		br = it.normalizeIndexBaseToArrayRef(a.X, br)

		return LValue{Kind: LVArrayElem, Base: br, Index: idx, SsaType: a.Type()}

	default:
		// обычный адрес: должен быть *symbolic.Ref
		e := it.resolveExpression(v)
		r, ok := e.(*symbolic.Ref)
		if !ok {
			panic(fmt.Sprintf("Addr is %T, want *Ref", e))
		}
		return LValue{Kind: LVDirectRef, Base: r, SsaType: v.Type()}
	}
}

func (it *Interpreter) load(lv LValue) symbolic.SymbolicExpression {
	switch lv.Kind {
	case LVDirectRef:
		return it.Heap.GetValue(lv.Base)

	case LVStructField:
		fr := it.Heap.StructFieldAddr(lv.Base, lv.Field)
		raw := it.Heap.GetValue(fr)

		// FieldAddr has type "*T", but we need T
		ft := lv.SsaType
		if p, ok := ft.Underlying().(*types.Pointer); ok {
			ft = p.Elem()
		}

		// 1) field is slice: []U / [][]U / ...
		if depth, base, ok := sliceDepthBase(ft); ok {
			// if field is nil in model => treat as base 0 (so IndexAddr still works)
			if raw.Type() == symbolic.NilType {
				raw = symbolic.NewIntConstant(0)
			}

			dim := depth - 1

			if typeName, fields, ok2 := structMeta(base); ok2 {
				if depth != 1 {
					panic("load struct field: multidimensional slice of structs is not supported")
				}
				r := symbolic.NewArrayRefFromExpr(raw, 0, symbolic.StructType)
				r.StructTypeName = typeName
				_ = fields
				return r
			}

			baseTy := goTypeToSym(base)
			return symbolic.NewArrayRefFromExpr(raw, dim, baseTy)
		}

		// 2) field is pointer: *U
		if p, ok := ft.Underlying().(*types.Pointer); ok {
			// nil pointer stays NilConstant
			if raw.Type() == symbolic.NilType {
				return raw
			}

			// pointer to struct
			if typeName, fields, ok2 := structMeta(p.Elem()); ok2 {
				_ = fields
				return symbolic.NewStructRefFromExpr(raw, typeName)
			}

			// pointer to primitive/string/float/bool
			elemTy := goTypeToSym(p.Elem())
			return symbolic.NewRefFromExpr(raw, elemTy)
		}

		// otherwise it's plain value
		return raw

	case LVArrayElem:
		eref := it.Heap.ArrayElemAddr(lv.Base, lv.Index)

		if eref.RefType == symbolic.StructType && eref.StructFieldIdx == -1 {
			return eref
		}

		return it.Heap.GetValue(eref)

	default:
		panic("unknown lvalue kind")
	}
}

func (it *Interpreter) store(lv LValue, val symbolic.SymbolicExpression) {
	switch lv.Kind {
	case LVDirectRef:
		it.Heap.AssignValue(lv.Base, val)

	case LVStructField:
		fr := it.Heap.StructFieldAddr(lv.Base, lv.Field)
		it.Heap.AssignValue(fr, val)

	case LVArrayElem:
		eref := it.Heap.ArrayElemAddr(lv.Base, lv.Index)
		it.Heap.AssignValue(eref, val)

	default:
		panic("unknown lvalue kind")
	}
}

// normalizeIndexBaseToArrayRef tries to "restore" array nature of base reference
// using SSA type information. This fixes cases when base was resolved as a scalar
// ref (e.g. bool) but is actually a slice/array.
func (it *Interpreter) normalizeIndexBaseToArrayRef(x ssa.Value, br *symbolic.Ref) *symbolic.Ref {
	if br == nil {
		panic("normalizeIndexBaseToArrayRef: nil base ref")
	}
	if br.RefType == symbolic.ArrayType {
		return br
	}

	// Determine base container type from SSA.
	// We prefer using existing helpers in this package: sliceDepthBase, structMeta, goTypeToSym.
	t := x.Type()

	// Slices: []T, [][]T, ...
	if depth, base, ok := sliceDepthBase(t); ok {
		// array dimension in your model:
		// []T => dim=0, [][]T => dim=1, [][][]T => dim=2, ...
		dim := depth - 1

		// []struct only allowed for 1D
		if typeName, fields, ok2 := structMeta(base); ok2 {
			if depth != 1 {
				panic("normalizeIndexBaseToArrayRef: multidimensional array of structs is not supported")
			}
			cp := *br
			cp.RefType = symbolic.ArrayType
			cp.ArrayDimension = 0
			cp.ArrayBaseType = symbolic.StructType
			cp.StructTypeName = typeName
			cp.StructFieldIdx = -1
			_ = fields // just to be explicit: memory type registration happens on Allocate/Param
			return &cp
		}

		baseTy := goTypeToSym(base)

		cp := *br
		cp.RefType = symbolic.ArrayType
		cp.ArrayDimension = dim
		cp.ArrayBaseType = baseTy
		cp.StructTypeName = ""
		cp.StructFieldIdx = -1
		return &cp
	}

	// Fixed-length arrays: [N]T (rare in your tests but harmless to support)
	// Also handles pointers to arrays if they appear via x.Type().
	tt := t
	for {
		if p, ok := tt.Underlying().(*types.Pointer); ok {
			tt = p.Elem()
			continue
		}
		break
	}
	if arr, ok := tt.Underlying().(*types.Array); ok {
		base := arr.Elem()

		if typeName, fields, ok2 := structMeta(base); ok2 {
			// treat as 1D array of structs
			cp := *br
			cp.RefType = symbolic.ArrayType
			cp.ArrayDimension = 0
			cp.ArrayBaseType = symbolic.StructType
			cp.StructTypeName = typeName
			cp.StructFieldIdx = -1
			_ = fields
			return &cp
		}

		baseTy := goTypeToSym(base)

		cp := *br
		cp.RefType = symbolic.ArrayType
		cp.ArrayDimension = 0
		cp.ArrayBaseType = baseTy
		cp.StructTypeName = ""
		cp.StructFieldIdx = -1
		return &cp
	}

	// If we got here, IndexAddr base is not an array-like type; keep as-is.
	// Better to panic to catch unexpected SSA shapes:
	panic(fmt.Sprintf("normalizeIndexBaseToArrayRef: base %s is not array-like (type=%s), baseRefType=%s",
		x.String(), x.Type().String(), br.RefType.String(),
	))
}

func (it *Interpreter) nilForType(t types.Type) symbolic.SymbolicExpression {
	zero := symbolic.NewIntConstant(0)

	// unwrap named
	for {
		if n, ok := t.(*types.Named); ok {
			t = n.Underlying()
			continue
		}
		break
	}

	// nil slice => keep as array handle 0 (so IndexAddr still works)
	if depth, base, ok := sliceDepthBase(t); ok {
		dim := depth - 1

		if typeName, fields, ok2 := structMeta(base); ok2 {
			if depth != 1 {
				panic("nilForType: multidimensional array of structs is not supported")
			}
			r := symbolic.NewArrayRefFromExpr(zero, 0, symbolic.StructType)
			r.StructTypeName = typeName
			_ = fields
			return r
		}

		baseTy := goTypeToSym(base)
		return symbolic.NewArrayRefFromExpr(zero, dim, baseTy)
	}

	// nil pointer => NilConstant (to be handled in binop)
	if _, ok := t.Underlying().(*types.Pointer); ok {
		return symbolic.NewNilConstant()
	}

	panic(fmt.Sprintf("nilForType: unsupported nil type: %s (%T)", t.String(), t))
}

func sanitizeName(s string) string {
	out := make([]rune, 0, len(s))
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			out = append(out, r)
		} else {
			out = append(out, '_')
		}
	}
	if len(out) == 0 {
		return "x"
	}
	return string(out)
}

func (it *Interpreter) cloneMapsOnly() Interpreter {
	cp := *it

	cp.Env = make(map[ssa.Value]symbolic.SymbolicExpression, len(it.Env))
	for k, v := range it.Env {
		cp.Env[k] = v
	}

	cp.LenCache = make(map[string]symbolic.SymbolicExpression, len(it.LenCache))
	for k, v := range it.LenCache {
		cp.LenCache[k] = v
	}

	cp.paramRef = make(map[ssa.Value]*symbolic.Ref, len(it.paramRef))
	for k, v := range it.paramRef {
		cp.paramRef[k] = v
	}

	if it.CallStack != nil {
		cp.CallStack = make([]CallStackFrame, len(it.CallStack))
		for i := range it.CallStack {
			cp.CallStack[i].Function = it.CallStack[i].Function
			cp.CallStack[i].ReturnValue = it.CallStack[i].ReturnValue

			if it.CallStack[i].LocalMemory != nil {
				cp.CallStack[i].LocalMemory = make(map[string]symbolic.SymbolicExpression, len(it.CallStack[i].LocalMemory))
				for k, v := range it.CallStack[i].LocalMemory {
					cp.CallStack[i].LocalMemory[k] = v
				}
			} else {
				cp.CallStack[i].LocalMemory = nil
			}
		}
	} else {
		cp.CallStack = nil
	}

	cp.Heap = cloneHeap(it.Heap)

	return cp
}

func cloneHeap(h memory.Memory) memory.Memory {
	if h == nil {
		return nil
	}
	sm, ok := h.(*memory.SymbolicMemory)
	if !ok {
		panic(fmt.Sprintf("cloneHeap: unsupported heap type %T", h))
	}
	return deepCopyAny[*memory.SymbolicMemory](sm)
}

func deepCopyAny[T any](in T) T {
	v := reflect.ValueOf(in)
	if !v.IsValid() {
		return in
	}
	cp := deepCopyValue(v, make(map[uintptr]reflect.Value))
	return cp.Interface().(T)
}

func deepCopyValue(v reflect.Value, memo map[uintptr]reflect.Value) reflect.Value {
	if !v.IsValid() {
		return v
	}

	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		ptr := v.Pointer()
		if got, ok := memo[ptr]; ok {
			return got
		}
		nv := reflect.New(v.Elem().Type())
		memo[ptr] = nv
		nv.Elem().Set(deepCopyValue(v.Elem(), memo))
		return nv

	case reflect.Interface:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		ev := deepCopyValue(v.Elem(), memo)
		nv := reflect.New(v.Type()).Elem()
		nv.Set(ev)
		return nv

	case reflect.Struct:
		nv := reflect.New(v.Type()).Elem()
		for i := 0; i < v.NumField(); i++ {
			sf := v.Field(i)
			cf := deepCopyValue(sf, memo)
			setValue(nv.Field(i), cf)
		}
		return nv

	case reflect.Slice:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		nv := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		for i := 0; i < v.Len(); i++ {
			nv.Index(i).Set(deepCopyValue(v.Index(i), memo))
		}
		return nv

	case reflect.Array:
		nv := reflect.New(v.Type()).Elem()
		for i := 0; i < v.Len(); i++ {
			nv.Index(i).Set(deepCopyValue(v.Index(i), memo))
		}
		return nv

	case reflect.Map:
		if v.IsNil() {
			return reflect.Zero(v.Type())
		}
		nv := reflect.MakeMapWithSize(v.Type(), v.Len())
		iter := v.MapRange()
		for iter.Next() {
			k := deepCopyValue(iter.Key(), memo)
			val := deepCopyValue(iter.Value(), memo)
			nv.SetMapIndex(k, val)
		}
		return nv

	default:
		return v
	}
}

func setValue(dst, src reflect.Value) {
	if dst.CanSet() {
		dst.Set(src)
		return
	}
	if dst.CanAddr() {
		reflect.NewAt(dst.Type(), unsafe.Pointer(dst.UnsafeAddr())).Elem().Set(src)
		return
	}
	panic(fmt.Sprintf("setValue: cannot set %s", dst.Type().String()))
}
