package symbolic

import (
	"strconv"
	"strings"
)

type Tensor struct {
	Name    string
	ValType ExpressionType

	Tensors []TensorExpression
}

func NewTensor(name string, valType ExpressionType) Tensor {
	return Tensor{
		Name:    name,
		ValType: valType,
		Tensors: make([]TensorExpression, 0),
	}
}

// SymbolicExpression

func (expr *Tensor) Type() ExpressionType {
	return TensorType
}

func (expr *Tensor) String() string {
	var b strings.Builder
	b.WriteString(expr.Name)
	b.WriteByte('[')

	for i, a := range expr.Tensors {
		if i > 0 {
			b.WriteByte(' ')
		}
		b.WriteByte('[')
		b.WriteString(strconv.Itoa(i))
		b.WriteString("]")
		b.WriteString(a.String())
	}

	b.WriteByte(']')
	return b.String()
}

func (expr *Tensor) ExprName() string {
	return expr.Name
}

func (expr *Tensor) Accept(visitor Visitor) interface{} {
	return visitor.VisitTensor(expr)
}

// TensorExpression Impl

func (expr *Tensor) Access(idx SymbolicExpression) SymbolicExpression {
	return &TensorAccess{Tensor: expr, Idx: idx}
}

func (expr *Tensor) BaseType() ExpressionType {
	return expr.ValType
}
