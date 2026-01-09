package memory

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	LookupRef(name string) *symbolic.Ref

	ParamRef(name string, tpe symbolic.ExpressionType) *symbolic.Ref

	ParamStructRef(name string, typeName string, fields []symbolic.ExpressionType) *symbolic.Ref

	ParamArrayRef(name string, tp symbolic.ExpressionType) *symbolic.Ref

	Allocate(tpe symbolic.ExpressionType) *symbolic.Ref

	AllocateStruct(name string, fields []symbolic.ExpressionType) *symbolic.Ref

	AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref

	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)

	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

	AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression)

	GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression

	AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression)

	GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression

	Constraints() []symbolic.SymbolicExpression
}

type SymbolicMemory struct {
	LastIdx        int
	ParamNameToRef map[string]*symbolic.Ref
	Params         []*symbolic.Ref
	Allocations    []*symbolic.Ref

	// Primitives
	IntSegment  *symbolic.Array
	BoolSegment *symbolic.Array

	// Structures
	LastTypeIdx  int
	RefToTypeIdx map[*symbolic.SymbolicVariable]int
	TypeToInt    map[string]int

	FieldsSegment map[int]map[int]*symbolic.Array

	// Arrays
	ArraysSegment map[int]*symbolic.Array
}

func NewSymbolicMemory() *SymbolicMemory {
	return &SymbolicMemory{
		LastIdx:        0,
		ParamNameToRef: make(map[string]*symbolic.Ref), //Cache to prevent duplicate same params in memory
		Params:         make([]*symbolic.Ref, 0),
		Allocations:    make([]*symbolic.Ref, 0),

		IntSegment:  symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment: symbolic.NewArray("BoolSegment", symbolic.BoolType),

		RefToTypeIdx: make(map[*symbolic.SymbolicVariable]int),
		TypeToInt:    make(map[string]int),

		FieldsSegment: make(map[int]map[int]*symbolic.Array),

		ArraysSegment: make(map[int]*symbolic.Array),
	}
}

func (mem *SymbolicMemory) LookupRef(name string) *symbolic.Ref {
	r, ok := mem.ParamNameToRef[name]
	if ok {
		// If Ref is found by name, return it, otherwirse go to create new
		return r
	}

	// Fallback
	return nil
}

func (mem *SymbolicMemory) ParamRef(name string, tpe symbolic.ExpressionType) *symbolic.Ref {
	r, ok := mem.ParamNameToRef[name]
	if ok {
		// If Ref is found by name, return it, otherwirse go to create new
		return r
	}

	switch tpe {
	case symbolic.IntType, symbolic.BoolType:
		r := symbolic.NewRef(mem.nextId(), tpe)
		mem.Params = append(mem.Params, r)
		mem.ParamNameToRef[name] = r
		return r

	default:
		panic(fmt.Sprintf("Allocate: Use another allocation method for %s", tpe.String()))
	}
}

func (mem *SymbolicMemory) ParamStructRef(name string, typeName string, fields []symbolic.ExpressionType) *symbolic.Ref {
	r, ok := mem.ParamNameToRef[name]
	if ok {
		// If Ref is found by name, return it, otherwirse go to create new
		return r
	}

	r = symbolic.NewRef(mem.nextId(), symbolic.StructType)
	mem.Params = append(mem.Params, r)

	typeIdx, ok := mem.TypeToInt[typeName]
	if !ok {
		typeIdx = int(mem.nextTypeId())
		mem.TypeToInt[typeName] = typeIdx

		mem.FieldsSegment[typeIdx] = make(map[int]*symbolic.Array)
		for fieldIdx, tpe := range fields {
			mem.FieldsSegment[typeIdx][fieldIdx] =
				symbolic.NewArray(fmt.Sprintf("%s_%d", typeName, fieldIdx), tpe)
		}
	}

	mem.RefToTypeIdx[r.Idx] = typeIdx
	mem.ParamNameToRef[name] = r

	return r
}

func (mem *SymbolicMemory) ParamArrayRef(name string, tp symbolic.ExpressionType) *symbolic.Ref {
	r, ok := mem.ParamNameToRef[name]
	if ok {
		// If Ref is found by name, return it, otherwirse go to create new
		return r
	}

	r = symbolic.NewRef(mem.nextId(), symbolic.ArrayType)
	mem.Params = append(mem.Params, r)
	mem.ParamNameToRef[name] = r

	return r
}

func (mem *SymbolicMemory) Allocate(tpe symbolic.ExpressionType) *symbolic.Ref {
	switch tpe {
	case symbolic.IntType, symbolic.BoolType:
		r := symbolic.NewRef(mem.nextId(), tpe)
		mem.Allocations = append(mem.Allocations, r)
		return r

	default:
		panic(fmt.Sprintf("Allocate: Use another allocation method for %s", tpe.String()))
	}
}

func (mem *SymbolicMemory) AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression) {
	switch ref.RefType {
	case symbolic.IntType:
		mem.IntSegment = symbolic.NewArrayStore(mem.IntSegment, ref.Idx, value)

	case symbolic.BoolType:
		mem.BoolSegment = symbolic.NewArrayStore(mem.BoolSegment, ref.Idx, value)

	default:
		panic(fmt.Sprintf("AssignValue: Unsupported operation for type %d", ref.RefType))
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

func (mem *SymbolicMemory) AllocateStruct(name string, fields []symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.StructType)
	mem.Allocations = append(mem.Allocations, r)

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

func (mem *SymbolicMemory) AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)
	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic("AssignToArray: ref is not array")
	}

	arr := mem.ArraysSegment[index]
	if arr == nil {
		// Если записи по одному и тому же
		// 	индексу для разных рефов равны, то можно поалиасить рефы
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", index), value.Type())
	}

	if arr.ValType != value.Type() {
		panic(fmt.Sprintf("AssignToArray: value type %s != array elem type %s",
			value.Type().String(), arr.ValType.String()))
	}

	mem.ArraysSegment[index] = symbolic.NewArrayStore(arr, ref.Idx, value)
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic("GetFromArray: ref is not array")
	}

	arr := mem.ArraysSegment[index]
	if arr == nil {
		arr = symbolic.NewArray(fmt.Sprintf("Array_%d", index), symbolic.IntType)
		mem.ArraysSegment[index] = arr
	}

	return symbolic.NewArrayAccess(arr, ref.Idx)
}

func (mem *SymbolicMemory) nextId() int64 {
	mem.LastIdx++
	return int64(mem.LastIdx - 1)
}

func (mem *SymbolicMemory) nextTypeId() int64 {
	mem.LastTypeIdx++
	return int64(mem.LastTypeIdx - 1)
}

func (mem *SymbolicMemory) Constraints() []symbolic.SymbolicExpression {
	ops := make([]symbolic.SymbolicExpression, 0)

	// params vs allocations
	for _, vp := range mem.Params {
		// fmt.Printf("vp : %s\n", vp.String())
		for _, va := range mem.Allocations {
			fmt.Printf("\tva : %s\n", va.String())
			ops = append(ops, symbolic.NewBinaryOperation(vp, va, symbolic.NE))
		}
	}

	// pairwise allocations (важно не добавлять va != va)
	for i := 0; i < len(mem.Allocations); i++ {
		for j := i + 1; j < len(mem.Allocations); j++ {
			ops = append(ops,
				symbolic.NewBinaryOperation(mem.Allocations[i], mem.Allocations[j], symbolic.NE),
			)
		}
	}

	return ops
}
