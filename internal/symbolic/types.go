package symbolic

// ExpressionType представляет тип символьного выражения
type ExpressionType int

const (
	IntType ExpressionType = iota
	BoolType
	StringType
	FloatType
	StructType
	RefType

	// Arrays
	ArrayType

	// Forbid type
	NONE
)

// String возвращает строковое представление типа
func (et ExpressionType) String() string {
	switch et {
	case IntType:
		return "int"
	case BoolType:
		return "bool"
	case StringType:
		return "string"
	case FloatType:
		return "float"
	case StructType:
		return "struct"
	case RefType:
		return "ref"
	case ArrayType:
		return "array"
	default:
		return "unknown"
	}
}

// Базовый интерфейс для всех символьных выражений
type SymbolicExpression interface {
	// Expression type
	Type() ExpressionType

	// String represenation
	String() string

	// Name
	ExprName() string

	// Visitor
	Accept(visitor Visitor) interface{}
}

// Многомерный массив
type ArrayExpression interface {
	SymbolicExpression

	// Ops
	Access(idx SymbolicExpression) SymbolicExpression
	Store(idx SymbolicExpression, val SymbolicExpression)

	// Info
	Dimension() int
	BaseType() ExpressionType
}
