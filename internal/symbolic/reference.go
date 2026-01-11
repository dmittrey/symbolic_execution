package symbolic

import "fmt"

type Ref struct {
	Idx     SymbolicExpression
	OwnType ExpressionType
	RefType ExpressionType

	// array refs
	ElemType ExpressionType

	// struct refs
	StructTypeIdx     int // какой именно struct type
	ElemStructTypeIdx int // если RefType==ArrayType && ElemType==StructType
}

func NewRef(idx int64, refType ExpressionType) *Ref {
	return &Ref{
		Idx:     NewSymbolicVariable(fmt.Sprintf("addr_%d", idx), IntType),
		OwnType: RefType,
		RefType: refType,

		ElemType:          -1,
		StructTypeIdx:     -1,
		ElemStructTypeIdx: -1,
	}
}
func NewRefExpr(idx SymbolicExpression, refType ExpressionType) *Ref {
	if idx == nil || idx.Type() != IntType {
		panic(fmt.Sprintf("NewRefExpr: idx must be int expression, got %T (%v)", idx, idx))
	}
	return &Ref{
		Idx:     idx,
		OwnType: RefType,
		RefType: refType,

		ElemType:          -1,
		StructTypeIdx:     -1,
		ElemStructTypeIdx: -1,
	}
}

func (r *Ref) Type() ExpressionType {
	return r.OwnType
}

func (r *Ref) String() string {
	return fmt.Sprintf("&(%s)", r.Idx.String())
}

func (r *Ref) ExprName() string {
	return ""
}

func (r *Ref) Accept(visitor Visitor) interface{} {
	panic("Forbid to translate pure ref")
}
