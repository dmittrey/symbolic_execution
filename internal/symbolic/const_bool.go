package symbolic

import "fmt"

// Булева константа
type BoolConstant struct {
	Value bool
}

func NewBoolConstant(value bool) *BoolConstant {
	return &BoolConstant{Value: value}
}

// SymbolicExpression

func (bc *BoolConstant) Type() ExpressionType {
	return BoolType
}

func (bc *BoolConstant) String() string {
	return fmt.Sprintf("%t", bc.Value)
}

func (bc *BoolConstant) ExprName() string {
	return ""
}

func (bc *BoolConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitBoolConstant(bc)
}
