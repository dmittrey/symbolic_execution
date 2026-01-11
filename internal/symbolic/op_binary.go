package symbolic

import "fmt"

// Бинарные операций
type BinaryOperation struct {
	Left     SymbolicExpression
	Right    SymbolicExpression
	Operator BinaryOperator
}

func NewBinaryOperation(left, right SymbolicExpression, op BinaryOperator) *BinaryOperation {
	lt, rt := left.Type(), right.Type()

	switch {
	case lt == IntType && rt == IntType && (op != IMPLIES):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == BoolType && rt == BoolType &&
		(op == EQ || op == NE || op == AND || op == OR || op == XOR || op == AND_NOT || op == IMPLIES):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == StringType && rt == StringType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == TensorType && rt == TensorType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == VectorType && rt == VectorType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == RefType && rt == RefType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left.(*Ref).Idx, Right: right.(*Ref).Idx, Operator: op}

	default:
		panic(fmt.Errorf("binary op %v not supported for %s(%v) and %s(%v)", op, left, lt, right, rt))
	}
}

// SymbolicExpression

func (bo *BinaryOperation) Type() ExpressionType {
	switch bo.Operator {
	case ADD, SUB, MUL, DIV, REM, XOR, SHL, SHR:
		return IntType
	case AND, OR, AND_NOT:
		if bo.Left.Type() == IntType {
			// Bitvec ops
			return IntType
		} else {
			// Bool ops
			return BoolType
		}
	case EQ, NE, LT, LE, GT, GE, IMPLIES:
		return BoolType
	default:
		panic("unknown operator")
	}
}

func (bo *BinaryOperation) String() string {
	return fmt.Sprintf("(%s %s %s)",
		bo.Left.String(),
		bo.Operator.String(),
		bo.Right.String(),
	)
}

func (bo *BinaryOperation) ExprName() string {
	return ""
}

func (bo *BinaryOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryOperation(bo)
}

// Операторы для бинарных выражений
type BinaryOperator int

const (
	// Арифметические операторы
	ADD BinaryOperator = iota
	SUB
	MUL
	DIV
	REM

	// Логические операторы
	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^
	IMPLIES // ->

	// Операторы сравнения
	EQ // равно
	NE // не равно
	LT // меньше
	LE // меньше или равно
	GT // больше
	GE // больше или равно
)

// String возвращает строковое представление оператора
func (op BinaryOperator) String() string {
	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case REM:
		return "%"
	case AND:
		return "&"
	case OR:
		return "|"
	case XOR:
		return "^"
	case SHL:
		return "<<"
	case SHR:
		return ">>"
	case AND_NOT:
		return "&^"
	case IMPLIES:
		return "=>"
	case EQ:
		return "=="
	case NE:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	default:
		panic("Impl me")
	}
}
