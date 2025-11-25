// Package symbolic определяет базовые типы символьных выражений
package symbolic

// ExpressionType представляет тип символьного выражения
type ExpressionType int

const (
	IntType ExpressionType = iota
	BoolType
	StringType
	ArrayType
	StructType
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
	case ArrayType:
		return "array"
	case StructType:
		return "struct"
	default:
		return "unknown"
	}
}
