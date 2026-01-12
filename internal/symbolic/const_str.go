package symbolic

import "fmt"

// Cтроковая константа
type StringConstant struct {
	Value string
}

func NewStringConstant(value string) *StringConstant {
	return &StringConstant{Value: value}
}

// SymbolicExpression

func (sc *StringConstant) Type() ExpressionType {
	return StringType
}

func (sc *StringConstant) String() string {
	return fmt.Sprintf("%q", sc.Value)
}

func (sc *StringConstant) ExprName() string {
	return ""
}

func (sc *StringConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitStringConstant(sc)
}
