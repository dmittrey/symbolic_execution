package symbolic

type SymbolicVariable struct {
	Name     string
	ExprType ExpressionType
}

func NewSymbolicVariable(name string, exprType ExpressionType) *SymbolicVariable {
	return &SymbolicVariable{
		Name:     name,
		ExprType: exprType,
	}
}

// SymbolicExpression

func (sv *SymbolicVariable) Type() ExpressionType {
	return sv.ExprType
}

func (sv *SymbolicVariable) String() string {
	return sv.Name
}

func (sv *SymbolicVariable) ExprName() string {
	return sv.Name
}

func (sv *SymbolicVariable) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariable(sv)
}
