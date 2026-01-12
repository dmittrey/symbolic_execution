package symbolic

import "fmt"

// Бинарные операций
type BinaryOperation struct {
	Left     SymbolicExpression
	Right    SymbolicExpression
	Operator BinaryOperator
}

func NewBinaryOperation(left, right SymbolicExpression, op BinaryOperator) SymbolicExpression {
	lt, rt := left.Type(), right.Type()

	switch {
	case lt == IntType && rt == IntType && op != IMPLIES:
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == FloatType && rt == FloatType && op != IMPLIES:
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == BoolType && rt == BoolType &&
		(op == EQ || op == NE || op == AND || op == OR || op == XOR || op == AND_NOT || op == IMPLIES):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == StringType && rt == StringType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == ArrayType && rt == ArrayType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == StructType && rt == StructType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left, Right: right, Operator: op}

	case lt == RefType && rt == RefType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left.(*Ref).Idx, Right: right.(*Ref).Idx, Operator: op}

	case lt == RefType && rt == NilType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: left.(*Ref).Idx, Right: NewRef(0, left.(*Ref).RefType), Operator: op}

	case lt == NilType && rt == RefType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: NewRef(0, right.(*Ref).RefType), Right: right.(*Ref).Idx, Operator: op}

	// Арифметика указателей
	case lt == RefType && rt == IntType &&
		(op == ADD || op == SUB):
		return left

	case lt == IntType && rt == RefType &&
		(op == ADD || op == SUB):
		return right

	case lt == NilType && rt == NilType &&
		(op == EQ || op == NE):
		return &BinaryOperation{Left: NewIntConstant(0), Right: NewIntConstant(0), Operator: op}

	default:
		panic(fmt.Errorf("binary op %v not supported for %s(%v) and %s(%v)", op, left, lt, right, rt))
	}
}

// SymbolicExpression

func (bo *BinaryOperation) Type() ExpressionType {
	// float арифметика
	if bo.Left.Type() == FloatType && bo.Right.Type() == FloatType {
		switch bo.Operator {
		case ADD, SUB, MUL, DIV:
			return FloatType
		case EQ, NE, LT, LE, GT, GE:
			return BoolType
		default:
			panic("unsupported float operator")
		}
	}

	switch bo.Operator {
	case ADD, SUB, MUL, DIV:
		if bo.Left.Type() == FloatType {
			return FloatType
		}
		if bo.Left.Type() == IntType {
			return IntType
		}
		return BoolType
	case REM, XOR, SHL, SHR:
		return IntType
	case AND, OR, AND_NOT:
		if bo.Left.Type() == IntType {
			return IntType
		}
		return BoolType
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
