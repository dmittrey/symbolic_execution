package symbolic

import "fmt"

type Ref struct {
	Idx     SymbolicExpression
	RefType ExpressionType

	// ArrayRef info
	ArrayDimension int
	ArrayBaseType  ExpressionType

	// StructRef info
}

func NewRef(idx int64, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: refType,
	}
}
func NewArrayRef(idx int64, dimension int, baseType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		RefType: ArrayType,

		ArrayDimension: dimension,
		ArrayBaseType:  baseType,
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
