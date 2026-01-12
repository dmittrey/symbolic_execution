package symbolic

import "fmt"

// Унарные операции
type UnaryOperation struct {
	Operand  SymbolicExpression
	Operator UnaryOperator
}

func NewUnaryOperation(operand SymbolicExpression, operator UnaryOperator) *UnaryOperation {
	switch {
	case operand.Type() == IntType && (operator == UXOR || operator == NEG):
		return &UnaryOperation{Operand: operand, Operator: operator}

	case operand.Type() == BoolType && operator == NOT:
		return &UnaryOperation{Operand: operand, Operator: operator}

	case operand.Type() == FloatType && operator == NEG:
		return &UnaryOperation{Operand: operand, Operator: operator}

	default:
		panic(fmt.Errorf("unary op %s not supported for %s(%s)", operator.String(), operand.String(), operand.Type().String()))
	}
}

// SymbolicExpression

func (uo *UnaryOperation) Type() ExpressionType {
	return uo.Operand.Type()
}

func (uo *UnaryOperation) String() string {
	return fmt.Sprintf("(%s%s)",
		uo.Operator.String(),
		uo.Operand.String(),
	)
}

func (uo *UnaryOperation) ExprName() string {
	return ""
}

func (uo *UnaryOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitUnaryOperation(uo)
}

// Операторы для унарных выражений
type UnaryOperator int

const (
	NOT UnaryOperator = iota
	NEG
	UXOR
)

func (op UnaryOperator) String() string {
	switch op {
	case NOT:
		return "!"
	case NEG:
		return "-"
	case UXOR:
		return "^"
	default:
		panic("Impl me")
	}
}
