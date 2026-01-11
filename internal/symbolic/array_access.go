package symbolic

type TensorAccess struct {
	Tensor TensorExpression
	Idx    SymbolicExpression
}

// SymbolicExpression

func (expr *TensorAccess) Type() ExpressionType {
	return expr.Tensor.BaseType()
}

func (expr *TensorAccess) String() string {
	return expr.Tensor.ExprName() + "[" + expr.Idx.String() + "]"
}

func (expr *TensorAccess) ExprName() string {
	// доступ по индексу — не "именованное" выражение
	return ""
}

func (expr *TensorAccess) Accept(visitor Visitor) interface{} {
	return visitor.VisitTensorAccess(expr)
}
