package symbolic

import "fmt"

type Ref struct {
	Idx     SymbolicExpression
	RefType ExpressionType

	// ArrayRef info
	ArrayDimension int
	ArrayBaseType  ExpressionType
	ArrayElemIdx   SymbolicExpression

	// StructRef info
	StructTypeName string
	StructFieldIdx int
}

func NewRef(idx int64, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: refType,

		ArrayDimension: -1,
		ArrayBaseType:  ForbiddenType,
		ArrayElemIdx:   nil,

		StructTypeName: "",
		StructFieldIdx: -1,
	}
}
func NewRefFromExpr(idx SymbolicExpression, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     idx,
		RefType: refType,

		ArrayDimension: -1,
		ArrayBaseType:  ForbiddenType,
		ArrayElemIdx:   nil,

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
		ArrayElemIdx:   nil,

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
		ArrayElemIdx:   nil,

		StructTypeName: "",
		StructFieldIdx: -1,
	}
}
func NewStructRef(idx int64, structTypeName string) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: StructType,

		ArrayDimension: -1,
		ArrayBaseType:  ForbiddenType,
		ArrayElemIdx:   nil,

		StructTypeName: structTypeName,
		StructFieldIdx: -1,
	}
}
func NewStructRefFromExpr(idx SymbolicExpression, structTypeName string) *Ref {
	return &Ref{
		Idx:     idx,
		RefType: StructType,

		ArrayDimension: -1,
		ArrayBaseType:  ForbiddenType,
		ArrayElemIdx:   nil,

		StructTypeName: structTypeName,
		StructFieldIdx: -1,
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
	return visitor.VisitRef(r)
}
