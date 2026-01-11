package memory

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	// Params
	Param(tp symbolic.ExpressionType) *symbolic.Ref
	ParamPrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref
	ParamComplexArray(baseType symbolic.ExpressionType) *symbolic.Ref

	// Primitives
	Allocate(tp symbolic.ExpressionType) *symbolic.Ref
	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)
	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

	// Arrays
	AllocatePrimitiveArray(baseType symbolic.ExpressionType) *symbolic.Ref
	AllocateComplexArray(dimension int, baseType symbolic.ExpressionType) *symbolic.Ref
	AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression)
	GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression

	// Aliasing constraints
	Constraints() []symbolic.SymbolicExpression
}

type SymbolicMemory struct {
	// Internals
	LastIdx int
	// Aliasing
	Params      []*symbolic.Ref
	Allocations []*symbolic.Ref

	// Primitive
	IntSegment    symbolic.Array //SymId -> BV
	BoolSegment   symbolic.Array //SymId -> Bool
	StringSegment symbolic.Array //SymId -> UF

	// Array
	ArraySegment symbolic.Array //SymId -> SymId
}

func NewSymbolicMemory() *SymbolicMemory {
	return &SymbolicMemory{
		LastIdx: 0,

		IntSegment:    symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment:   symbolic.NewArray("BoolSegment", symbolic.BoolType),
		StringSegment: symbolic.NewArray("StringSegment", symbolic.StringType),

		ArraySegment: symbolic.NewArray("ArraySegment", symbolic.IntType),
	}
}

func (mem *SymbolicMemory) Param(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), tp)
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

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic(fmt.Sprintf("%s %s", ref.RefType, symbolic.ArrayType))
	}

	// 1) Взять id у ref
	refKey := ref.Idx
	refDim := ref.ArrayDimension

	// 2) Раскручиваем косвенность через ArraySegment
	// Bool[]   - 1
	// Bool[][] - 2
	for refDim != 0 {
		refKey = mem.ArraySegment.Access(refKey)
		refDim--
	}

	// 3) Пойти в нужный сегмент, получить значение по адресу
	addr := mem.effAddr(refKey, index)
	switch ref.ArrayBaseType {
	case symbolic.IntType:
		mem.IntSegment.Store(addr, value)
	case symbolic.BoolType:
		mem.BoolSegment.Store(addr, value)
	case symbolic.StringType:
		mem.StringSegment.Store(addr, value)
	default:
		panic("unknown base type")
	}
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic(fmt.Sprintf("%s %s", ref.RefType, symbolic.ArrayType))
	}

	// 1) Взять id у ref
	refKey := ref.Idx
	refDim := ref.ArrayDimension

	// 2) Раскручиваем косвенность через ArraySegment
	// Bool[]   - 1
	// Bool[][] - 2
	for refDim != 0 {
		refKey = mem.ArraySegment.Access(refKey)
		refDim--
	}

	// 3) Пойти в нужный сегмент, получить значение по адресу
	addr := mem.effAddr(refKey, index)
	switch ref.ArrayBaseType {
	case symbolic.IntType:
		return mem.IntSegment.Access(addr)
	case symbolic.BoolType:
		return mem.BoolSegment.Access(addr)
	case symbolic.StringType:
		return mem.StringSegment.Access(addr)
	default:
		panic("unknown base type")
	}
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
	return []symbolic.SymbolicExpression{
		&mem.IntSegment,
		&mem.BoolSegment,
		&mem.StringSegment,
		&mem.ArraySegment,
	}
}

func (mem *SymbolicMemory) effAddr(base symbolic.SymbolicExpression, index symbolic.SymbolicExpression) symbolic.SymbolicExpression {
	// base + index
	return symbolic.NewBinaryOperation(base, index, symbolic.ADD)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastIdx++
	return int64(mem.LastIdx - 1)
}
