package symbolic

import "fmt"

type Ref struct {
	Idx     SymbolicExpression
	RefType ExpressionType

	// ArrayRef info
	ArrayDimension int
	ArrayBaseType  ExpressionType

	// StructRef info
	StructTypeName string
	StructFieldIdx int
}

func NewRef(idx int64, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: refType,

		ArrayDimension: -1,
		ArrayBaseType:  NONE,

		StructTypeName: "",
		StructFieldIdx: -1,
	}
}
func NewArrayRef(idx int64, dimension int, baseType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: ArrayType,

		ArrayDimension: dimension,
		ArrayBaseType:  baseType,

		StructTypeName: "",
		StructFieldIdx: -1,
	}
}
func NewArrayRefFromExpr(idx SymbolicExpression, dimension int, baseType ExpressionType) *Ref {
	return &Ref{
		Idx:     idx,
		RefType: ArrayType,

		ArrayDimension: dimension,
		ArrayBaseType:  baseType,

		StructTypeName: "",
		StructFieldIdx: -1,
	}
}
func NewStructRef(idx int64, typeName string, fieldIdx int) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: StructType,

		ArrayDimension: -1,
		ArrayBaseType:  NONE,

		StructTypeName: typeName,
		StructFieldIdx: fieldIdx,
	}
}

func (r *Ref) Type() ExpressionType {
	return RefType
}

func (r *Ref) String() string {
	return fmt.Sprintf("&(%s)", r.Idx.String())
}

func (r *Ref) ExprName() string {
	return ""
}

func (r *Ref) Accept(visitor Visitor) interface{} {
	panic("Forbid to translate pure ref")
}
