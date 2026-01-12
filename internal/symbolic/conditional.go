package symbolic

import "fmt"

// ConditionalExpression (тернарный оператор: condition ? true_expr : false_expr)
type ConditionalExpression struct {
	IfOperand   SymbolicExpression
	ThenOperand SymbolicExpression
	ElseOperand SymbolicExpression
}

func NewConditionalExpression(ifOperand SymbolicExpression, thenOperand SymbolicExpression, elseOperand SymbolicExpression) *ConditionalExpression {
	return &ConditionalExpression{
		IfOperand:   ifOperand,
		ThenOperand: thenOperand,
		ElseOperand: elseOperand,
	}
}

// SymbolicExpression

func (ce *ConditionalExpression) Type() ExpressionType {
	if ce.IfOperand.Type() != BoolType {
		panic("condition of ?: must be bool")
	}
	if ce.ThenOperand.Type() != ce.ElseOperand.Type() {
		panic("then/else types must match")
	}
	return ce.ThenOperand.Type()
}

func (ce *ConditionalExpression) String() string {
	return fmt.Sprintf("(%s ? %s : %s)",
		ce.IfOperand.String(), ce.ThenOperand.String(), ce.ElseOperand.String())
}

func (ce *ConditionalExpression) ExprName() string {
	return ""
}

func (ce *ConditionalExpression) Accept(visitor Visitor) interface{} {
	return visitor.VisitConditionalExpression(ce)
}
