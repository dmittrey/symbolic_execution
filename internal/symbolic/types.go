package symbolic

// ExpressionType представляет тип символьного выражения
type ExpressionType int

const (
	// Primitives
	IntType ExpressionType = iota
	BoolType
	StringType
	FloatType

	// Reference
	RefType

	// Struct
	StructType

	// Arrays
	ArrayType

	// Nil
	NilType

	// Not initialized
	ForbiddenType
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
	case RefType:
		return "ref"
	case StructType:
		return "struct"
	case ArrayType:
		return "array"
	case NilType:
		return "Nil"
	case ForbiddenType:
		return "FORBID"
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
}
