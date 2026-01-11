package symbolic

import (
	"fmt"
	"strings"
)

type ArrayOp int

const (
	OpStore ArrayOp = iota
	OpSelect
)

type Op struct {
	Kind ArrayOp
	Idx  SymbolicExpression
	Val  SymbolicExpression // store: значение; select: переменная-результат
}

type Array struct {
	Name     string
	ElemType ExpressionType

	// ops — последовательность операций (store/select)
	Ops []Op
}

func NewArray(name string, elemType ExpressionType) Array {
	return Array{
		Name:     name,
		ElemType: elemType,
		Ops:      make([]Op, 0),
	}
}

// SymbolicExpression

func (t *Array) Type() ExpressionType { return ArrayType }

func (t *Array) ExprName() string { return t.Name }

func (t *Array) String() string {
	var b strings.Builder
	b.WriteString(t.Name)

	for _, op := range t.Ops {
		switch op.Kind {
		case OpStore:
			b.WriteString(".store(")
			b.WriteString(op.Idx.String())
			b.WriteString(", ")
			b.WriteString(op.Val.String())
			b.WriteByte(')')

		case OpSelect:
			b.WriteString(".select(")
			b.WriteString(op.Idx.String())
			b.WriteString(", ")
			b.WriteString(op.Val.String())
			b.WriteByte(')')

		default:
			b.WriteString(".<unknown-op>")
		}
	}

	return b.String()
}

func (t *Array) Accept(visitor Visitor) interface{} {
	return visitor.VisitArray(t)
}

// ArrayExpression

func (a *Array) Access(idx SymbolicExpression) SymbolicExpression {
	res := NewSymbolicVariable(fmt.Sprintf("%s[%s]", a.Name, idx.String()), a.ElemType)
	a.Ops = append(a.Ops, Op{Kind: OpSelect, Idx: idx, Val: res})
	return res
}

func (a *Array) Store(idx, val SymbolicExpression) {
	a.Ops = append(a.Ops, Op{Kind: OpStore, Idx: idx, Val: val})
}
