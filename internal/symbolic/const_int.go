package symbolic

import "fmt"

// Целочисленная константа
type IntConstant struct {
	Value int64
}

func NewIntConstant(value int64) *IntConstant {
	return &IntConstant{Value: value}
}

// SymbolicExpression

func (ic *IntConstant) Type() ExpressionType {
	return IntType
}

func (ic *IntConstant) String() string {
	return fmt.Sprintf("%d", ic.Value)
}

func (ic *IntConstant) ExprName() string {
	return ""
}

func (ic *IntConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitIntConstant(ic)
}
