package symbolic

import "fmt"

type FloatConstant struct {
	Value float64
}

func NewFloatConstant(v float64) *FloatConstant {
	return &FloatConstant{Value: v}
}

func (fc *FloatConstant) Type() ExpressionType {
	return FloatType
}
func (fc *FloatConstant) String() string {
	return fmt.Sprintf("%g", fc.Value)
}
func (fc *FloatConstant) ExprName() string {
	return ""
}
func (fc *FloatConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitFloatConstant(fc)
}
