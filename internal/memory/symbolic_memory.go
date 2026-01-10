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

	AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression)

	GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression

	Constraints() []symbolic.SymbolicExpression

	ParamArrayStructRef(name, typeName string, fields []symbolic.ExpressionType) *symbolic.Ref
	AllocateArrayStruct(typeName string, fields []symbolic.ExpressionType) *symbolic.Ref
}

type SymbolicMemory struct {
	LastIdx        int
	ParamNameToRef map[string]*symbolic.Ref
	Params         []*symbolic.Ref
	Allocations    []*symbolic.Ref

	// Primitives
	IntSegment    *symbolic.Array
	BoolSegment   *symbolic.Array
	StringSegment *symbolic.Array

	// Structures
	LastTypeIdx int
	TypeToInt   map[string]int

	FieldsSegment map[int]map[int]*symbolic.Array

	// Arrays
	ArraysSegment map[string]*symbolic.Array
}

func NewSymbolicMemory() *SymbolicMemory {
	return &SymbolicMemory{
		LastIdx:        0,
		ParamNameToRef: make(map[string]*symbolic.Ref), //Cache to prevent duplicate same params in memory
		Params:         make([]*symbolic.Ref, 0),
		Allocations:    make([]*symbolic.Ref, 0),

		IntSegment:    symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment:   symbolic.NewArray("BoolSegment", symbolic.BoolType),
		StringSegment: symbolic.NewArray("StringSegment", symbolic.StringType),

		TypeToInt: make(map[string]int),

		FieldsSegment: make(map[int]map[int]*symbolic.Array),

		ArraysSegment: make(map[string]*symbolic.Array),
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
	case symbolic.IntType, symbolic.BoolType, symbolic.StringType:
		r := symbolic.NewRef(mem.nextId(), tpe)
		mem.Params = append(mem.Params, r)
		mem.ParamNameToRef[name] = r
		return r

	default:
		panic(fmt.Sprintf("Allocate: Use another allocation method for %s", tpe.String()))
	}
}

func (mem *SymbolicMemory) ParamStructRef(name, typeName string, fields []symbolic.ExpressionType) *symbolic.Ref {
	if r, ok := mem.ParamNameToRef[name]; ok {
		return r
	}

	r := symbolic.NewRef(mem.nextId(), symbolic.StructType)
	r.StructTypeIdx = mem.ensureStructType(typeName, fields)

	mem.Params = append(mem.Params, r)
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
	r.ElemType = tp

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

	case symbolic.StringType:
		mem.StringSegment = symbolic.NewArrayStore(mem.StringSegment, ref.Idx, value)

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

	case symbolic.StringType:
		return symbolic.NewArrayAccess(mem.StringSegment, ref.Idx)

	default:
		panic("AssignValue: Unsupported operation for type")
	}
}

func (mem *SymbolicMemory) AllocateStruct(name string, fields []symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.StructType)
	r.StructTypeIdx = mem.ensureStructType(name, fields)

	mem.Allocations = append(mem.Allocations, r)
	return r
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	typeIdx := ref.StructTypeIdx
	if typeIdx == -1 {
		// если вдруг используешь 0 как валидный — сделай -1 или отдельный флаг
		panic("AssignField: StructTypeIdx is not set on ref")
	}

	curArr := mem.FieldsSegment[typeIdx][fieldIdx]
	mem.FieldsSegment[typeIdx][fieldIdx] = symbolic.NewArrayStore(curArr, ref.Idx, value)
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	typeIdx := ref.StructTypeIdx
	if typeIdx == -1 {
		panic("GetFieldValue: StructTypeIdx is not set on ref")
	}
	curArr := mem.FieldsSegment[typeIdx][fieldIdx]
	return symbolic.NewArrayAccess(curArr, ref.Idx)
}

func (mem *SymbolicMemory) AllocateArray(tp symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)
	r.ElemType = tp

	mem.Allocations = append(mem.Allocations, r)

	return r
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index symbolic.SymbolicExpression, value symbolic.SymbolicExpression) {
	if ref.RefType != symbolic.ArrayType {
		panic("AssignToArray: ref is not array")
	}
	if index.Type() != symbolic.IntType {
		panic(fmt.Sprintf("AssignToArray: index must be int, got %s", index.Type().String()))
	}

	key := addrKey(ref)
	elemTy := ref.ElemType

	// ---- CASE: array of struct pointers ----
	if elemTy == symbolic.StructType {
		vr, ok := value.(*symbolic.Ref)
		if !ok || vr.RefType != symbolic.StructType {
			panic(fmt.Sprintf("AssignToArray: want *Ref(StructType), got %T (%s)", value, value.String()))
		}
		// массив хранит АДРЕСА (int)
		arr := mem.ArraysSegment[key]
		if arr == nil {
			arr = symbolic.NewArray(fmt.Sprintf("Arr_%s", key), symbolic.IntType)
		}
		mem.ArraysSegment[key] = symbolic.NewArrayStore(arr, index, vr.Idx)
		return
	}

	// ---- old cases: primitives ----
	if elemTy != symbolic.IntType && elemTy != symbolic.BoolType && elemTy != symbolic.StringType {
		panic(fmt.Sprintf("AssignToArray: unknown elem type %s", elemTy.String()))
	}
	if value.Type() != elemTy {
		panic(fmt.Sprintf("AssignToArray: value type %s != elem type %s", value.Type().String(), elemTy.String()))
	}

	arr := mem.ArraysSegment[key]
	if arr == nil {
		arr = symbolic.NewArray(fmt.Sprintf("Arr_%s", key), elemTy)
	}
	mem.ArraysSegment[key] = symbolic.NewArrayStore(arr, index, value)
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index symbolic.SymbolicExpression) symbolic.SymbolicExpression {
	if ref.RefType != symbolic.ArrayType {
		panic("GetFromArray: ref is not array")
	}
	if index.Type() != symbolic.IntType {
		panic(fmt.Sprintf("GetFromArray: index must be int, got %s", index.Type().String()))
	}

	key := addrKey(ref)
	elemTy := ref.ElemType

	// ---- CASE: array of struct pointers ----
	if elemTy == symbolic.StructType {
		arr := mem.ArraysSegment[key]
		if arr == nil {
			arr = symbolic.NewArray(fmt.Sprintf("Arr_%s", key), symbolic.IntType)
			mem.ArraysSegment[key] = arr
		}

		addrExpr := symbolic.NewArrayAccess(arr, index) // IntType
		r := symbolic.NewRefExpr(addrExpr, symbolic.StructType)
		r.StructTypeIdx = ref.ElemStructTypeIdx
		if r.StructTypeIdx == -1 {
			panic("GetFromArray: ElemStructTypeIdx is not set on array ref")
		}
		return r
	}

	// ---- old cases: primitives ----
	if elemTy != symbolic.IntType && elemTy != symbolic.BoolType && elemTy != symbolic.StringType {
		panic(fmt.Sprintf("GetFromArray: unknown elem type %s", elemTy.String()))
	}

	arr := mem.ArraysSegment[key]
	if arr == nil {
		arr = symbolic.NewArray(fmt.Sprintf("Arr_%s", key), elemTy)
		mem.ArraysSegment[key] = arr
	}
	return symbolic.NewArrayAccess(arr, index)
}

func (mem *SymbolicMemory) ParamArrayStructRef(name, typeName string, fields []symbolic.ExpressionType) *symbolic.Ref {
	if r, ok := mem.ParamNameToRef[name]; ok {
		return r
	}

	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)
	r.ElemType = symbolic.StructType
	r.ElemStructTypeIdx = mem.ensureStructType(typeName, fields)

	mem.Params = append(mem.Params, r)
	mem.ParamNameToRef[name] = r
	return r
}

func (mem *SymbolicMemory) AllocateArrayStruct(typeName string, fields []symbolic.ExpressionType) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), symbolic.ArrayType)
	r.ElemType = symbolic.StructType
	r.ElemStructTypeIdx = mem.ensureStructType(typeName, fields)

	mem.Allocations = append(mem.Allocations, r)
	return r
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

func addrKey(ref *symbolic.Ref) string {
	// ref.Idx — SymbolicExpression, но у реальных refs это всегда SymbolicVariable
	if v, ok := ref.Idx.(*symbolic.SymbolicVariable); ok {
		return v.Name
	}
	// если вдруг когда-нибудь ref.Idx станет чем-то сложнее — ты это сразу увидишь
	panic(fmt.Sprintf("ref.Idx is not *SymbolicVariable: %T (%s)", ref.Idx, ref.Idx.String()))
}

func (mem *SymbolicMemory) ensureStructType(typeName string, fields []symbolic.ExpressionType) int {
	typeIdx, ok := mem.TypeToInt[typeName]
	if ok {
		return typeIdx
	}

	typeIdx = int(mem.nextTypeId())
	mem.TypeToInt[typeName] = typeIdx

	mem.FieldsSegment[typeIdx] = make(map[int]*symbolic.Array)
	for fieldIdx, tpe := range fields {
		mem.FieldsSegment[typeIdx][fieldIdx] =
			symbolic.NewArray(fmt.Sprintf("%s_%d", typeName, fieldIdx), tpe)
	}
	return typeIdx
}
