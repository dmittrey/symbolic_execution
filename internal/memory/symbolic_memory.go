package memory

import (
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	Allocate(tpe symbolic.ExpressionType) *symbolic.Ref

	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)

	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

	AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression)

	GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression

	AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression)

	GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression
}

type SymbolicMemory struct {
	LastIdx         int
	NotAliasingRefs []*symbolic.Ref

	RefSegment *symbolic.Array

	IntSegment  *symbolic.Array
	BoolSegment *symbolic.Array
}

func NewSymbolicMemory() SymbolicMemory {
	return SymbolicMemory{
		LastIdx:         0,
		NotAliasingRefs: make([]*symbolic.Ref, 0),

		RefSegment: symbolic.NewArray("RefSegment", symbolic.IntType),

		IntSegment:  symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment: symbolic.NewArray("BoolSegment", symbolic.BoolType),
	}
}

func (mem *SymbolicMemory) Allocate(tpe symbolic.ExpressionType) *symbolic.Ref {
	switch tpe {
	case symbolic.IntType, symbolic.BoolType:
		r := symbolic.NewRef(mem.nextId(), tpe)

		// Prevent aliasing for primitives
		rIdx := mem.getRefIdx(r)
		mem.RefSegment = symbolic.NewArrayStore(mem.RefSegment, rIdx, rIdx)

		return r
	default:
		//TODO implement me
		panic("implement me")
	}
}

// Ref lvalue -> SymExpr (primitive variant)
func (mem *SymbolicMemory) AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression) {
	switch ref.RefType {
	case symbolic.IntType:
		IntSegmentIdx := symbolic.NewArrayAccess(mem.RefSegment, symbolic.NewIntConstant(ref.Idx))

		mem.IntSegment = symbolic.NewArrayStore(mem.IntSegment, IntSegmentIdx, value)
	case symbolic.BoolType:
		BoolSegmentIdx := symbolic.NewArrayAccess(mem.RefSegment, symbolic.NewIntConstant(ref.Idx))

		mem.BoolSegment = symbolic.NewArrayStore(mem.BoolSegment, BoolSegmentIdx, value)
	default:
		//TODO implement me
		panic("implement me")
	}
}

func (mem *SymbolicMemory) GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression {
	switch ref.RefType {
	case symbolic.IntType:
		return symbolic.NewArrayAccess(
			mem.IntSegment,
			symbolic.NewArrayAccess(mem.RefSegment, symbolic.NewIntConstant(ref.Idx)))
	case symbolic.BoolType:
		return symbolic.NewArrayAccess(
			mem.BoolSegment,
			symbolic.NewArrayAccess(mem.RefSegment, symbolic.NewIntConstant(ref.Idx)))
	default:
		//TODO implement me
		panic("implement me")
	}
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) getRefIdx(ref *symbolic.Ref) *symbolic.IntConstant {
	return symbolic.NewIntConstant(ref.Idx)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastIdx++
	return int64(mem.LastIdx - 1)
}
