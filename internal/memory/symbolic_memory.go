package memory

import (
	"fmt"
	"strconv"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	// Alloc
	Allocate(tp symbolic.ExpressionType, isAliasing bool) *symbolic.Ref
	AllocateArray(dimension int, baseType symbolic.ExpressionType, isAliasing bool) *symbolic.Ref
	AllocateStruct(structTypeName string, fieldTypes []symbolic.ExpressionType, isAliasing bool) *symbolic.Ref
	AllocateStructsArray(dimension int, structName string, fieldTypes []symbolic.ExpressionType, isAliasing bool) *symbolic.Ref

	// Address
	StructFieldAddr(ref *symbolic.Ref, fieldIdx int) *symbolic.Ref
	ArrayElemAddr(ref *symbolic.Ref, elemIdx symbolic.SymbolicExpression) *symbolic.Ref

	// Ops
	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)
	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

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
	FloatSegment  symbolic.Array //SymId -> Real

	// Array
	ArraySegment symbolic.Array //SymId -> SymId

	// Struct
	StructSegments map[string]map[int]*symbolic.Array //TypeName -> FieldIdx -> SymId -> FieldSort
}

func NewSymbolicMemory() *SymbolicMemory {
	return &SymbolicMemory{
		LastRefIdx: 0,

		Params:      make([]*symbolic.Ref, 0),
		Allocations: make([]*symbolic.Ref, 0),

		IntSegment:    symbolic.NewArray("IntSegment", symbolic.IntType),
		BoolSegment:   symbolic.NewArray("BoolSegment", symbolic.BoolType),
		StringSegment: symbolic.NewArray("StringSegment", symbolic.StringType),
		FloatSegment:  symbolic.NewArray("FloatSegment", symbolic.FloatType),

		ArraySegment: symbolic.NewArray("ArraySegment", symbolic.IntType),

		StructSegments: make(map[string]map[int]*symbolic.Array),
	}
}

func (mem *SymbolicMemory) Allocate(tp symbolic.ExpressionType, isAliasing bool) *symbolic.Ref {
	r := symbolic.NewRef(mem.nextId(), tp)
	mem.updateAliasing(r, isAliasing)

	return r
}

func (mem *SymbolicMemory) AllocateArray(dimension int, baseType symbolic.ExpressionType, isAliasing bool) *symbolic.Ref {
	r := symbolic.NewArrayRef(mem.nextId(), dimension, baseType)
	mem.updateAliasing(r, isAliasing)

	return r
}

func (mem *SymbolicMemory) AllocateStruct(structTypeName string, fieldTypes []symbolic.ExpressionType, isAliasing bool) *symbolic.Ref {
	// Prepare field segments
	mem.registerStructType(structTypeName, fieldTypes)

	r := symbolic.NewStructRef(mem.nextId(), structTypeName)
	mem.updateAliasing(r, isAliasing)

	return r
}

func (mem *SymbolicMemory) AllocateStructsArray(dimension int, structTypeName string, fieldTypes []symbolic.ExpressionType, isAliasing bool) *symbolic.Ref {
	// Prepare field segments
	mem.registerStructType(structTypeName, fieldTypes)

	r := symbolic.NewArrayRef(mem.nextId(), dimension, symbolic.StructType)
	r.StructTypeName = structTypeName
	mem.updateAliasing(r, isAliasing)

	return r
}

func (mem *SymbolicMemory) StructFieldAddr(ref *symbolic.Ref, fieldIdx int) *symbolic.Ref {
	if ref == nil {
		panic("StructFieldAddr: nil ref")
	}
	if ref.RefType != symbolic.StructType {
		panic("StructFieldAddr: ref is not struct")
	}
	if ref.StructTypeName == "" {
		panic("StructFieldAddr: empty struct type name")
	}
	if mem.StructSegments[ref.StructTypeName] == nil {
		panic("StructFieldAddr: struct type is not registered")
	}
	if mem.StructSegments[ref.StructTypeName][fieldIdx] == nil {
		panic("StructFieldAddr: field segment is not registered")
	}

	cp := *ref
	cp.StructFieldIdx = fieldIdx
	return &cp
}

func (mem *SymbolicMemory) ArrayElemAddr(ref *symbolic.Ref, elemIdx symbolic.SymbolicExpression) *symbolic.Ref {
	if ref == nil {
		panic("ArrayElemAddr: nil ref")
	}
	if ref.RefType != symbolic.ArrayType {
		panic("ArrayElemAddr: ref is not array")
	}
	if elemIdx == nil {
		panic("ArrayElemAddr: nil elemIdx")
	}

	slotAddr := mem.effAddr(ref.Idx, elemIdx)

	// Terminal dimension
	if ref.ArrayDimension == 0 {
		switch ref.ArrayBaseType {
		case symbolic.IntType, symbolic.BoolType, symbolic.StringType, symbolic.FloatType:
			return symbolic.NewRefFromExpr(slotAddr, ref.ArrayBaseType)

		case symbolic.StructType:
			// Take struct handle
			handle := mem.ArraySegment.Access(slotAddr)
			return symbolic.NewStructRefFromExpr(handle, ref.StructTypeName)

		default:
			panic("ArrayElemAddr: unknown base type")
		}
	}

	// Not terminal dimension
	cp := *ref
	cp.ArrayDimension -= 1
	cp.ArrayElemIdx = elemIdx
	return &cp
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
	case symbolic.FloatType:
		mem.FloatSegment.Store(refKey, value)
	case symbolic.StructType:
		if ref.StructFieldIdx == -1 {
			mem.assignWholeStruct(ref, value)
			return
		}

		fieldSegment := mem.StructSegments[ref.StructTypeName][ref.StructFieldIdx]
		if fieldSegment == nil {
			panic("AssignValue: struct field segment is nil")
		}
		fieldSegment.Store(refKey, value)
	case symbolic.ArrayType:
		if ref.ArrayElemIdx == nil {
			panic("AssignValue: array ref without ArrayElemIdx (not an element address)")
		}

		slotAddr := mem.effAddr(ref.Idx, ref.ArrayElemIdx)

		// тут мы храним handle подмассива
		if vref, ok := value.(*symbolic.Ref); ok {
			if vref.RefType != symbolic.ArrayType {
				panic("AssignValue(Array): expected array ref as value for non-last dimension")
			}
			mem.ArraySegment.Store(slotAddr, vref.Idx) // slot -> handle
			return
		}

		// если ты допускаешь raw handle
		mem.ArraySegment.Store(slotAddr, value)

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
	case symbolic.FloatType:
		return mem.FloatSegment.Access(refKey)
	case symbolic.StructType:
		if ref.StructFieldIdx == -1 {
			panic("GetValue: can't read whole struct; take field addr and read field")
		}
		return mem.StructSegments[ref.StructTypeName][ref.StructFieldIdx].Access(refKey)
	case symbolic.ArrayType:
		if ref.ArrayElemIdx == nil {
			panic("GetValue: array ref without ArrayElemIdx (not an element address)")
		}

		slotAddr := mem.effAddr(ref.Idx, ref.ArrayElemIdx)
		childHandle := mem.ArraySegment.Access(slotAddr)

		r := symbolic.NewArrayRefFromExpr(childHandle, ref.ArrayDimension, ref.ArrayBaseType)
		r.StructTypeName = ref.StructTypeName
		return r

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
	out := []symbolic.SymbolicExpression{
		&mem.IntSegment,
		&mem.BoolSegment,
		&mem.StringSegment,
		&mem.FloatSegment,
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

// Update Params and Allocations collections to generate new constraints(distinct/not)
// when asserting final interpreter states
func (mem *SymbolicMemory) updateAliasing(ref *symbolic.Ref, isAliasing bool) {
	if isAliasing {
		// May aliasing
		mem.Params = append(mem.Params, ref)
	} else {
		// Not aliasing
		mem.Allocations = append(mem.Allocations, ref)
	}
}

// *dst = *src
func (mem *SymbolicMemory) assignWholeStruct(dst *symbolic.Ref, srcExpr symbolic.SymbolicExpression) {
	src, ok := srcExpr.(*symbolic.Ref)
	if !ok || src.RefType != symbolic.StructType {
		panic("AssignValue: assigning non-struct into whole struct")
	}
	if src.StructTypeName != dst.StructTypeName {
		panic("AssignValue: struct type mismatch in whole-struct assignment")
	}

	fields := mem.StructSegments[dst.StructTypeName]
	if fields == nil {
		panic("AssignValue: struct type not registered: " + dst.StructTypeName)
	}

	// dst.field[i] = src.field[i]
	for fieldIdx := range fields {
		dstField := mem.StructFieldAddr(dst, fieldIdx)
		srcField := mem.StructFieldAddr(src, fieldIdx)
		mem.AssignValue(dstField, mem.GetValue(srcField))
	}
}
