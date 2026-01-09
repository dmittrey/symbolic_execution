package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"

	"golang.org/x/tools/go/ssa"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	ReturnResults []symbolic.SymbolicExpression
	Heap          memory.Memory
	Env           map[ssa.Value]symbolic.SymbolicExpression
}

func NewInterpreter(analyser *Analyser, pathCond symbolic.SymbolicExpression) Interpreter {
	return Interpreter{
		CallStack:     make([]CallStackFrame, 0),
		Analyser:      analyser,
		PathCondition: pathCond,
		Heap:          memory.NewSymbolicMemory(),
		Env:           make(map[ssa.Value]symbolic.SymbolicExpression),
	}
}

type CallStackFrame struct {
	Function    *ssa.Function
	LocalMemory map[string]symbolic.SymbolicExpression
	ReturnValue symbolic.SymbolicExpression
}

func (interpreter *Interpreter) interpretDynamically(block *ssa.BasicBlock, instrIdx int, pred *ssa.BasicBlock) []*Item {
	if block == nil || instrIdx < 0 || instrIdx >= len(block.Instrs) {
		return nil
	}

	instr := block.Instrs[instrIdx]
	fmt.Printf("interpretDynamically: %s\n", instr.String())

	if instr == nil {
		fmt.Println("interpretDynamically with nil!")
		return nil
	}

	switch i := instr.(type) {
	case *ssa.Return:
		interpreter.ReturnResults = make([]symbolic.SymbolicExpression, len(i.Results))
		for idx, val := range i.Results {
			interpreter.ReturnResults[idx] = interpreter.resolveExpression(val)
		}

		return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}
	case *ssa.Panic:
		// No results on panic
		interpreter.ReturnResults = make([]symbolic.SymbolicExpression, 0)

		return []*Item{{value: *interpreter, block: nil, instrIdx: -1}}
	case *ssa.Select:
		panic("Support me!")
	case *ssa.Store:
		panic("Support me!")
	case *ssa.Jump:
		out := make([]*Item, 0, len(block.Succs))
		for _, s := range block.Succs {
			out = append(out, &Item{value: *interpreter, block: s, instrIdx: 0, pred: block})
		}
		return out
	case *ssa.If:
		cond := interpreter.resolveExpression(i.Cond)

		out := make([]*Item, 0, 2)

		// then -> Succs[0]
		if len(i.Block().Succs) >= 1 {
			thenState := *interpreter // копия
			thenState.PathCondition = symbolic.NewBinaryOperation(
				thenState.PathCondition,
				cond,
				symbolic.AND,
			)
			out = append(out, &Item{value: thenState, block: block.Succs[0], instrIdx: 0, pred: block})
		}
		if len(i.Block().Succs) >= 2 {
			elseState := *interpreter // копия
			notCond := symbolic.NewBinaryOperation(
				cond,
				symbolic.NewBoolConstant(true), //TODO Add unary
				symbolic.AND_NOT,
			)
			elseState.PathCondition = symbolic.NewBinaryOperation(
				elseState.PathCondition,
				notCond,
				symbolic.AND,
			)
			out = append(out, &Item{value: elseState, block: block.Succs[1], instrIdx: 0, pred: block})
		}
		return out
	case *ssa.Phi:
		edge, ok := pickPhiEdge(i, block, pred)
		if !ok {
			panic("phi: can't pick edge (pred is nil or not found)")
		}
		phiVal := interpreter.resolveExpression(edge)

		interpreter.Env[i] = phiVal

		return nextInSameBlock(*interpreter, block, instrIdx, pred)

	// Instructions that also implement Value interface will resolve later
	case *ssa.BinOp, *ssa.UnOp:
		return nextInSameBlock(*interpreter, block, instrIdx, pred)
	default:
		panic(fmt.Sprintf("Support %T!", i))
	}
}

func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {
	if value == nil {
		fmt.Println("resolveExpression with nil!")
		return nil
	}

	fmt.Printf("resolveExpression: %s\n", value.String())

	if e, ok := interpreter.Env[value]; ok {
		return e
	}

	switch v := value.(type) {
	// A Parameter represents an input parameter of a function.
	case *ssa.Parameter:
		ref := interpreter.Heap.ParamRef(v.Name(), goTypeToSym(v.Type()))
		return interpreter.Heap.GetValue(ref)

	// IntConstant, BoolConstant, StringConstant
	case *ssa.Const:
		if v.Value == nil {
			panic("unsupported nil value")
		}

		switch v.Value.Kind() {
		case constant.Bool:
			return symbolic.NewBoolConstant(constant.BoolVal(v.Value))

		case constant.Int:
			i, ok := constant.Int64Val(v.Value)
			if !ok {
				panic(fmt.Sprintf("big int const not supported: %s", v.Value.String()))
			}
			return symbolic.NewIntConstant(i)

		case constant.String:
			return symbolic.NewStringConstant(constant.StringVal(v.Value))

		default:
			panic(fmt.Sprintf("unsupported value %T", v.Value.Kind()))
		}

	// X Op Y
	/*
		One of:
		ADD SUB MUL QUO REM          + - * / %
		AND OR XOR SHL SHR AND_NOT   & | ^ << >> &^
		EQL NEQ LSS LEQ GTR GEQ      == != < <= < >=
	*/
	case *ssa.BinOp:
		l := interpreter.resolveExpression(v.X)
		r := interpreter.resolveExpression(v.Y)

		if sop, ok := tokenToBinOp(v.Op); ok {
			return symbolic.NewBinaryOperation(l, r, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	// Op X
	/*
		One of:
		NOT SUB ARROW MUL XOR ! - <- * ^
	*/
	case *ssa.UnOp:
		operand := interpreter.resolveExpression(v.X)

		if sop, ok := tokenToUnOp(v.Op); ok {
			return symbolic.NewUnaryOperation(operand, sop)
		}

		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))

	// The Call instruction represents a function or method call.
	case *ssa.Call:
		panic("Implement me HW5!")

	default:
		// fallback: просто символ по имени SSA value
		panic("Support me!")
	}
}

func tokenToBinOp(op token.Token) (symbolic.BinaryOperator, bool) {
	switch op {
	case token.ADD:
		return symbolic.ADD, true
	case token.SUB:
		return symbolic.SUB, true
	case token.MUL:
		return symbolic.MUL, true
	case token.QUO:
		return symbolic.DIV, true
	case token.REM:
		return symbolic.REM, true

	case token.AND:
		return symbolic.AND, true
	case token.OR:
		return symbolic.OR, true
	case token.XOR:
		return symbolic.XOR, true
	case token.SHL:
		return symbolic.SHL, true
	case token.SHR:
		return symbolic.SHR, true
	case token.AND_NOT:
		return symbolic.AND_NOT, true

	case token.EQL:
		return symbolic.EQ, true
	case token.NEQ:
		return symbolic.NE, true
	case token.LSS:
		return symbolic.LT, true
	case token.LEQ:
		return symbolic.LE, true
	case token.GTR:
		return symbolic.GT, true
	case token.GEQ:
		return symbolic.GE, true
	default:
		panic("Not impl!")
	}
}

func tokenToUnOp(op token.Token) (symbolic.UnaryOperator, bool) {
	switch op {
	case token.NOT:
		return symbolic.NOT, true
	case token.SUB:
		return symbolic.NEG, true
	case token.XOR:
		return symbolic.UXOR, true
	default:
		panic("Not impl!")
	}
}

func goTypeToSym(t types.Type) symbolic.ExpressionType {
	if t == nil {
		return symbolic.IntType // fallback
	}
	switch tt := t.Underlying().(type) {
	case *types.Basic:
		switch {
		case tt.Info()&types.IsBoolean != 0:
			return symbolic.BoolType
		case tt.Info()&types.IsInteger != 0:
			return symbolic.IntType
		case tt.Info()&types.IsString != 0:
			return symbolic.StringType
		}
	}
	// упрощение: всё неизвестное как Int
	return symbolic.IntType
}

func nextInSameBlock(st Interpreter, block *ssa.BasicBlock, instrIdx int, pred *ssa.BasicBlock) []*Item {
	nxt := instrIdx + 1
	if nxt < len(block.Instrs) {
		return []*Item{{value: st, block: block, instrIdx: nxt, pred: pred}}
	}
	out := make([]*Item, 0, len(block.Succs))
	for _, s := range block.Succs {
		out = append(out, &Item{value: st, block: s, instrIdx: 0, pred: block})
	}
	return out
}

func pickPhiEdge(phi *ssa.Phi, block *ssa.BasicBlock, pred *ssa.BasicBlock) (ssa.Value, bool) {
	if pred == nil {
		return nil, false
	}
	for idx, p := range block.Preds {
		if p == pred {
			if idx < 0 || idx >= len(phi.Edges) {
				return nil, false
			}
			return phi.Edges[idx], true
		}
	}
	return nil, false
}
