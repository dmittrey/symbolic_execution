package memory

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
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
	LastIdx int

	// Primitives
	IntSegment  *symbolic.Array
	BoolSegment *symbolic.Array

	// Structures
	LastTypeIdx  int
	RefToTypeIdx map[*symbolic.SymbolicVariable]int
	TypeToInt    map[string]int

	FieldsSegment map[int]map[int]*symbolic.Array

	// Arrays
	ArraysSegment map[*symbolic.SymbolicVariable]*symbolic.Array
}

func NewSymbolicMemory() SymbolicMemory {
	return SymbolicMemory{
		LastIdx: 0,

		IntSegment:  symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment: symbolic.NewArray("BoolSegment", symbolic.BoolType),

		RefToTypeIdx: make(map[*symbolic.SymbolicVariable]int),
		TypeToInt:    make(map[string]int),

		FieldsSegment: make(map[int]map[int]*symbolic.Array),

		ArraysSegment: make(map[*symbolic.SymbolicVariable]*symbolic.Array),
	}
}

func (mem *SymbolicMemory) Allocate(tpe symbolic.ExpressionType) *symbolic.Ref {
	switch tpe {
	case symbolic.IntType, symbolic.BoolType:
		return symbolic.NewRef(mem.nextId(), tpe)

	default:
		panic(fmt.Sprintf("Allocate: Use another allocation method for %s", tpe.String()))
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

	mem.RefToTypeIdx[r.Idx] = typeIdx

	return r
}

func (mem *SymbolicMemory) AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)

	mem.ArraysSegment[r.Idx] =
		symbolic.NewArray(fmt.Sprintf("Array_%d", r.Idx), tp)

	return r
}

func (mem *SymbolicMemory) AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression) {
	switch ref.RefType {
	case symbolic.IntType:
		mem.IntSegment = symbolic.NewArrayStore(mem.IntSegment, ref.Idx, value)

	case symbolic.BoolType:
		mem.BoolSegment = symbolic.NewArrayStore(mem.BoolSegment, ref.Idx, value)

	default:
		panic("AssignValue: Unsupported operation for type")
	}
}

func (mem *SymbolicMemory) GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression {
	switch ref.RefType {
	case symbolic.IntType:
		return symbolic.NewArrayAccess(mem.IntSegment, ref.Idx)

	case symbolic.BoolType:
		return symbolic.NewArrayAccess(mem.BoolSegment, ref.Idx)

	default:
		panic("AssignValue: Unsupported operation for type")
	}
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	TypeIdx := mem.RefToTypeIdx[ref.Idx]

	FieldToArray := mem.FieldsSegment[TypeIdx]
	curArr := FieldToArray[fieldIdx]
	updArr := symbolic.NewArrayStore(curArr, ref.Idx, value)

	mem.FieldsSegment[TypeIdx][fieldIdx] = updArr
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	TypeIdx := mem.RefToTypeIdx[ref.Idx]

	FieldToArray := mem.FieldsSegment[TypeIdx]
	curArr := FieldToArray[fieldIdx]

	return symbolic.NewArrayAccess(curArr, ref.Idx)
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic("AssignToArray: ref is not array")
	}

	arr := mem.ArraysSegment[ref.Idx]
	if arr == nil {
		// если вдруг не инициализирован – создаём, тип берём из value
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", ref.Idx), value.Type())
	}

	if arr.ValType != value.Type() {
		panic(fmt.Sprintf("AssignToArray: value type %s != array elem type %s",
			value.Type().String(), arr.ValType.String()))
	}

	idxExpr := symbolic.NewIntConstant(int64(index))
	mem.ArraysSegment[ref.Idx] = symbolic.NewArrayStore(arr, idxExpr, value)
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic("GetFromArray: ref is not array")
	}

	arr := mem.ArraysSegment[ref.Idx]
	if arr == nil {
		// Лениво инициализируем массив int'ов
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", ref.Idx), symbolic.IntType)
		mem.ArraysSegment[ref.Idx] = arr
	}

	idxExpr := symbolic.NewIntConstant(int64(index))
	return symbolic.NewArrayAccess(arr, idxExpr)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastIdx++
	return int64(mem.LastIdx - 1)
}

func (mem *SymbolicMemory) nextTypeId() int64 {
	mem.LastTypeIdx++
	return int64(mem.LastTypeIdx - 1)
}
