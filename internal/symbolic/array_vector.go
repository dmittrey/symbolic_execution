package symbolic

import "fmt"

type vectorItem struct {
	Idx SymbolicExpression
	Val SymbolicExpression
}

type Vector struct {
	Name    string
	ValType ExpressionType
	Stores  []vectorItem
}

func NewVector(name string, valType ExpressionType) Vector {
	return Vector{
		Name:    name,
		ValType: valType,
	}
}

// SymbolicExpression

func (expr *Vector) Type() ExpressionType {
	return VectorType
}

func (expr *Vector) String() string {
	return expr.Name + "[]"
}

func (expr *Vector) ExprName() string {
	return expr.Name
}

func (expr *Vector) Accept(visitor Visitor) interface{} {
	return visitor.VisitVector(expr)
}

// TensorExpression

func (array *Vector) Access(idx SymbolicExpression) SymbolicExpression {
	return &TensorAccess{Tensor: array, Idx: idx}
}

func (expr *Vector) BaseType() ExpressionType {
	return expr.ValType
}

// VectorExpression

func (array *Vector) Store(idx SymbolicExpression, val SymbolicExpression) {
	if val.Type() != array.ValType {
		panic(fmt.Sprintf("Array value must be %s", array.ValType.String()))
	}

	array.Stores = append(array.Stores, vectorItem{idx, val})
}
