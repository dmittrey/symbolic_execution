package symbolic

import "strings"

type ArrayOp int

const (
	OpStore ArrayOp = iota
	OpSelect
)

type Op struct {
	Kind ArrayOp
	Idx  SymbolicExpression
	Val  SymbolicExpression // только для store; для select = nil
}

type Array struct {
	Name string

	// ops — последовательность операций (store/select)
	Ops []Op
}

func NewArray(name string, valType ExpressionType) *Array {
	return &Array{
		Name: name,
		Ops:  make([]Op, 0),
	}
}

// SymbolicExpression

func (t *Array) Type() ExpressionType { return TensorType }

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

func (t *Array) Access(idx SymbolicExpression) {
	t.Ops = append(t.Ops, Op{Kind: OpSelect, Idx: idx})
}

func (t *Array) Store(idx, val SymbolicExpression) {
	t.Ops = append(t.Ops, Op{Kind: OpStore, Idx: idx, Val: val})
}
