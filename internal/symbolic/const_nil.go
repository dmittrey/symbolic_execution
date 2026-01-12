package symbolic

// Булева константа
type NilConstant struct{}

func NewNilConstant() *NilConstant {
	return &NilConstant{}
}

// SymbolicExpression

func (nc *NilConstant) Type() ExpressionType {
	return NilType
}

func (nc *NilConstant) String() string {
	return "nil"
}

func (nc *NilConstant) ExprName() string {
	return "nil"
}

func (nc *NilConstant) Accept(visitor Visitor) interface{} {
	panic("NilConstant: Need to interpret in binop like pair arg sort!")
}
