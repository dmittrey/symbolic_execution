package memory

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	Param(tpe symbolic.ExpressionType) *symbolic.Ref

	Allocate(tpe symbolic.ExpressionType) *symbolic.Ref

	AllocateStruct(name string, fields []symbolic.ExpressionType) *symbolic.Ref

	AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref

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

	LastTypeIdx  int
	RefToTypeIdx map[int]int
	TypeToInt    map[string]int

	FieldsSegment map[int]map[int]*symbolic.Array

	ArraysSegment map[int]*symbolic.Array
}

func NewSymbolicMemory() SymbolicMemory {
	return SymbolicMemory{
		LastIdx:         0,
		NotAliasingRefs: make([]*symbolic.Ref, 0),

		RefSegment: symbolic.NewArray("RefSegment", symbolic.IntType),

		IntSegment:  symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment: symbolic.NewArray("BoolSegment", symbolic.BoolType),

		RefToTypeIdx:  make(map[int]int),
		TypeToInt:     make(map[string]int),
		FieldsSegment: make(map[int]map[int]*symbolic.Array),

		ArraysSegment: make(map[int]*symbolic.Array),
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
		panic(fmt.Sprintf("Allocate: Use another allocation method for %s", tpe.String()))
	}
}

func (mem *SymbolicMemory) ParamRef(tpe symbolic.ExpressionType) *symbolic.Ref {
	return symbolic.NewRef(mem.nextId(), tpe)
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
		panic("AssignValue: Unsupported operation for type")
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
		panic("AssignValue: Unsupported operation for type")
	}
}

func (mem *SymbolicMemory) AllocateStruct(name string, fields []symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.StructType)

	typeIdx, ok := mem.TypeToInt[name]
	if !ok {
		typeIdx = int(mem.nextTypeId())
		mem.TypeToInt[name] = typeIdx

		mem.FieldsSegment[typeIdx] = make(map[int]*symbolic.Array)
		for fieldIdx, tpe := range fields {
			mem.FieldsSegment[typeIdx][fieldIdx] =
				symbolic.NewArray(fmt.Sprintf("%s_%d", name, fieldIdx), tpe)
		}
	}

	mem.RefToTypeIdx[int(r.Idx)] = typeIdx

	return r
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	TypeIdx := mem.RefToTypeIdx[int(ref.Idx)]

	FieldToArray := mem.FieldsSegment[TypeIdx]
	curArr := FieldToArray[fieldIdx]
	updArr := symbolic.NewArrayStore(curArr, symbolic.NewIntConstant(ref.Idx), value)

	mem.FieldsSegment[TypeIdx][fieldIdx] = updArr
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	TypeIdx := mem.RefToTypeIdx[int(ref.Idx)]

	FieldToArray := mem.FieldsSegment[TypeIdx]
	curArr := FieldToArray[fieldIdx]

	return symbolic.NewArrayAccess(curArr, symbolic.NewIntConstant(ref.Idx))
}

func (mem *SymbolicMemory) AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)

	mem.ArraysSegment[int(r.Idx)] =
		symbolic.NewArray(fmt.Sprintf("Array_%d", r.Idx), tp)

	return r
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic("AssignToArray: ref is not array")
	}

	arr := mem.ArraysSegment[int(ref.Idx)]
	if arr == nil {
		// если вдруг не инициализирован – создаём, тип берём из value
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", ref.Idx), value.Type())
	}

	if arr.ValType != value.Type() {
		panic(fmt.Sprintf("AssignToArray: value type %s != array elem type %s",
			value.Type().String(), arr.ValType.String()))
	}

	idxExpr := symbolic.NewIntConstant(int64(index))
	mem.ArraysSegment[int(ref.Idx)] = symbolic.NewArrayStore(arr, idxExpr, value)
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic("GetFromArray: ref is not array")
	}

	arr := mem.ArraysSegment[int(ref.Idx)]
	if arr == nil {
		// Лениво инициализируем массив int'ов
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", ref.Idx), symbolic.IntType)
		mem.ArraysSegment[int(ref.Idx)] = arr
	}

	idxExpr := symbolic.NewIntConstant(int64(index))
	return symbolic.NewArrayAccess(arr, idxExpr)
}

func (mem *SymbolicMemory) getRefIdx(ref *symbolic.Ref) *symbolic.IntConstant {
	return symbolic.NewIntConstant(ref.Idx)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastIdx++
	return int64(mem.LastIdx - 1)
}

func (mem *SymbolicMemory) nextTypeId() int64 {
	mem.LastTypeIdx++
	return int64(mem.LastTypeIdx - 1)
}
