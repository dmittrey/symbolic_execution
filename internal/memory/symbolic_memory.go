package memory

import (
	"fmt"
	"strconv"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	// Params
	Param(tp symbolic.ExpressionType) *symbolic.Ref
	ParamStruct(structTypeName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref
	ParamPrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref
	ParamComplexArray(baseType symbolic.ExpressionType) *symbolic.Ref
	ParamStructsArray(dimension int, structName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref

	// Primitives
	Allocate(tp symbolic.ExpressionType) *symbolic.Ref
	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)
	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

	// Arrays
	AllocatePrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref
	AllocateComplexArray(dimension int, baseType symbolic.ExpressionType) *symbolic.Ref
	AllocateStructsArray(dimension int, structName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref
	AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression)
	GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression

	// Struct
	AllocateStruct(structTypeName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref
	FieldAddress(ref *symbolic.Ref, fieldIdx int) *symbolic.Ref

	// Aliasing constraints
	Constraints() []symbolic.SymbolicExpression
}

type SymbolicMemory struct {
	// Internals
	LastRefIdx int

	// Aliasing
	Params      []*symbolic.Ref
	Allocations []*symbolic.Ref

	// Primitive
	IntSegment    symbolic.Array //SymId -> BV
	BoolSegment   symbolic.Array //SymId -> Bool
	StringSegment symbolic.Array //SymId -> UF

	// Array
	ArraySegment symbolic.Array //SymId -> SymId

	// Struct
	StructSegments map[string]map[int]*symbolic.Array
}

func NewSymbolicMemory() *SymbolicMemory {
	return &SymbolicMemory{
		LastRefIdx: 0,

		Params:      make([]*symbolic.Ref, 0),
		Allocations: make([]*symbolic.Ref, 0),

		IntSegment:    symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment:   symbolic.NewArray("BoolSegment", symbolic.BoolType),
		StringSegment: symbolic.NewArray("StringSegment", symbolic.StringType),

		ArraySegment: symbolic.NewArray("ArraySegment", symbolic.IntType),

		StructSegments: make(map[string]map[int]*symbolic.Array),
	}
}

func (mem *SymbolicMemory) Param(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), tp)
	mem.Params = append(mem.Params, r)

	return r
}

func (mem *SymbolicMemory) ParamStruct(structTypeName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref {
	mem.registerStructType(structTypeName, fieldTypes)

	r := symbolic.NewRef(mem.nextId(), symbolic.StructType)
	r.StructTypeName = structTypeName
	mem.Params = append(mem.Params, r)

	return r
}

func (mem *SymbolicMemory) ParamPrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewArrayRef(mem.nextId(), 0, baseType)
	mem.Params = append(mem.Params, r)

	return r
}

func (mem *SymbolicMemory) ParamComplexArray(dimension int, baseType symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewArrayRef(mem.nextId(), dimension, baseType)
	mem.Params = append(mem.Params, r)

	return r
}

func (mem *SymbolicMemory) ParamStructsArray(dimension int, structTypeName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref {
	mem.registerStructType(structTypeName, fieldTypes)

	r := symbolic.NewArrayRef(mem.nextId(), dimension, symbolic.StructType)
	r.StructTypeName = structTypeName
	mem.Params = append(mem.Params, r)

	return r
}

func (mem *SymbolicMemory) Allocate(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), tp)
	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression) {
	refKey := ref.Idx

	switch ref.RefType {
	case symbolic.IntType:
		mem.IntSegment.Store(refKey, value)
	case symbolic.BoolType:
		mem.BoolSegment.Store(refKey, value)
	case symbolic.StringType:
		mem.StringSegment.Store(refKey, value)
	case symbolic.StructType:
		(mem.StructSegments[ref.StructTypeName][ref.StructFieldIdx]).Store(refKey, value)
	default:
		panic("unknown ref type")
	}
}

func (mem *SymbolicMemory) GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression {
	refKey := ref.Idx

	switch ref.RefType {
	case symbolic.IntType:
		return mem.IntSegment.Access(refKey)
	case symbolic.BoolType:
		return mem.BoolSegment.Access(refKey)
	case symbolic.StringType:
		return mem.StringSegment.Access(refKey)
	case symbolic.StructType:
		return (mem.StructSegments[ref.StructTypeName][ref.StructFieldIdx]).Access(refKey)
	default:
		panic("unknown base type")
	}
}

func (mem *SymbolicMemory) AllocatePrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewArrayRef(mem.nextId(), 0, baseType)
	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AllocateComplexArray(dimension int, baseType symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewArrayRef(mem.nextId(), dimension, baseType)
	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AllocateStructsArray(dimension int, structTypeName string, fieldTypes []symbolic.ExpressionType) *symbolic.Ref {
	// Prepare memory internals
	mem.registerStructType(structTypeName, fieldTypes)

	r := symbolic.NewArrayRef(mem.nextId(), dimension, symbolic.StructType)
	r.StructTypeName = structTypeName
	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic(fmt.Sprintf("%s %s", ref.RefType, symbolic.ArrayType))
	}

	slotAddr := mem.effAddr(ref.Idx, index)

	// Если не последнее измерение, то ref[index] — это handle (адрес) подмассива.
	if ref.ArrayDimension > 0 {
		// Разрешим 2 формы:
		// 1) value — это *Ref на массив (пишем его Idx как адрес)
		if vref, ok := value.(*symbolic.Ref); ok {
			if vref.RefType != symbolic.ArrayType {
				panic("AssignToArray: expected array ref as value for non-last dimension")
			}
			mem.ArraySegment.Store(slotAddr, vref.Idx)
			return
		}

		// 2) или value — сразу символьный адрес (BV), если ты так моделируешь handle
		// (оставь, если тебе это нужно; иначе лучше panic)
		mem.ArraySegment.Store(slotAddr, value)
		return
	}

	// Последнее измерение — пишем скаляр
	switch ref.ArrayBaseType {
	case symbolic.IntType:
		mem.IntSegment.Store(slotAddr, value)
	case symbolic.BoolType:
		mem.BoolSegment.Store(slotAddr, value)
	case symbolic.StringType:
		mem.StringSegment.Store(slotAddr, value)
	case symbolic.StructType:
		panic("AssignToArray: assigning whole struct is not supported; assign fields instead")
	default:
		panic("unknown base type")
	}
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic(fmt.Sprintf("%s %s", ref.RefType, symbolic.ArrayType))
	}

	slotAddr := mem.effAddr(ref.Idx, index)

	// Если это ещё НЕ последнее измерение — возвращаем handle подмассива
	if ref.ArrayDimension > 0 {
		childBase := mem.ArraySegment.Access(slotAddr)
		r := symbolic.NewArrayRefFromExpr(childBase, ref.ArrayDimension-1, ref.ArrayBaseType)
		// если массив структур — протащим тип
		r.StructTypeName = ref.StructTypeName
		return r
	}

	// Последнее измерение — читаем скаляр
	switch ref.ArrayBaseType {
	case symbolic.IntType:
		return mem.IntSegment.Access(slotAddr)
	case symbolic.BoolType:
		return mem.BoolSegment.Access(slotAddr)
	case symbolic.StringType:
		return mem.StringSegment.Access(slotAddr)
	case symbolic.StructType:
		// Потом в FieldAddress доуточним
		objAddr := mem.ArraySegment.Access(slotAddr) // BV address of struct object
		return &symbolic.Ref{
			Idx:            objAddr,
			RefType:        symbolic.StructType,
			StructTypeName: ref.StructTypeName,
		}
	default:
		panic("unknown base type")
	}
}

func (mem *SymbolicMemory) FieldAddress(ref *symbolic.Ref, fieldIdx int) *symbolic.Ref {
	if mem.StructSegments[ref.StructTypeName] == nil {
		panic("Not found type in indexing types")
	}

	ref.StructFieldIdx = fieldIdx
	return ref
}

func (mem *SymbolicMemory) Constraints() []symbolic.SymbolicExpression {
	ops := make([]symbolic.SymbolicExpression, 0)

	// params vs allocations
	for _, vp := range mem.Params {
		for _, va := range mem.Allocations {
			ops = append(ops, symbolic.NewBinaryOperation(vp.Idx, va.Idx, symbolic.NE))
		}
	}

	// pairwise allocations (важно не добавлять va != va)
	for i := 0; i < len(mem.Allocations); i++ {
		for j := i + 1; j < len(mem.Allocations); j++ {
			ops = append(ops,
				symbolic.NewBinaryOperation(mem.Allocations[i].Idx, mem.Allocations[j].Idx, symbolic.NE),
			)
		}
	}

	return append(ops, mem.segmentConstraints()...)
}

func (mem *SymbolicMemory) segmentConstraints() []symbolic.SymbolicExpression {
	out := []symbolic.SymbolicExpression{
		&mem.IntSegment,
		&mem.BoolSegment,
		&mem.StringSegment,
		&mem.ArraySegment,
	}
	for _, fields := range mem.StructSegments {
		for _, seg := range fields {
			out = append(out, seg) // seg is *symbolic.Array
		}
	}
	return out
}

func (mem *SymbolicMemory) effAddr(base symbolic.SymbolicExpression, index symbolic.SymbolicExpression) symbolic.SymbolicExpression {
	// base + index
	return symbolic.NewBinaryOperation(base, index, symbolic.ADD)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastRefIdx++
	return int64(mem.LastRefIdx - 1)
}

func (mem *SymbolicMemory) registerStructType(structTypeName string, fieldTypes []symbolic.ExpressionType) {
	// already registered => do nothing
	if mem.StructSegments[structTypeName] != nil {
		return
	}

	mem.StructSegments[structTypeName] = make(map[int]*symbolic.Array)
	for idx, ft := range fieldTypes {
		arr := symbolic.NewArray(fmt.Sprintf("%s_%s", structTypeName, strconv.Itoa(idx)), ft)
		mem.StructSegments[structTypeName][idx] = &arr
	}
}
