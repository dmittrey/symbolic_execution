package symbolic

// Visitor интерфейс для обхода символьных выражений (Visitor Pattern)
type Visitor interface {
	VisitVariable(expr *SymbolicVariable) interface{}
	VisitIntConstant(expr *IntConstant) interface{}
	VisitBoolConstant(expr *BoolConstant) interface{}
	VisitStringConstant(*StringConstant) interface{}
	VisitFloatConstant(expr *FloatConstant) interface{}
	VisitBinaryOperation(expr *BinaryOperation) interface{}
	VisitUnaryOperation(expr *UnaryOperation) interface{}
	VisitConditionalExpression(expr *ConditionalExpression) interface{}
	VisitRef(r *Ref) interface{}
	VisitArray(expr *Array) interface{}
}
