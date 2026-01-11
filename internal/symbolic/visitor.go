package symbolic

// Visitor интерфейс для обхода символьных выражений (Visitor Pattern)
type Visitor interface {
	VisitVariable(expr *SymbolicVariable) interface{}
	VisitIntConstant(expr *IntConstant) interface{}
	VisitBoolConstant(expr *BoolConstant) interface{}
	VisitBinaryOperation(expr *BinaryOperation) interface{}
	VisitUnaryOperation(expr *UnaryOperation) interface{}
	VisitConditionalExpression(expr *ConditionalExpression) interface{}
	VisitStringConstant(*StringConstant) interface{}

	// Arrays
	VisitVector(expr *Vector) interface{}
	VisitTensor(expr *Tensor) interface{}
	VisitTensorAccess(expr *TensorAccess) interface{}
}
