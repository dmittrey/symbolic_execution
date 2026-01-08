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
	Heap          memory.Memory
}

type CallStackFrame struct {
	Function    *ssa.Function
	LocalMemory map[string]symbolic.SymbolicExpression
	ReturnValue symbolic.SymbolicExpression
}

func (interpreter *Interpreter) interpretDynamically(instr ssa.Instruction) []Interpreter {
	if instr == nil {
		fmt.Println("interpretDynamically with nil!")
		return nil
	}

	switch i := instr.(type) {
	case *ssa.Return:
		panic("Support me!")
	case *ssa.Panic:
		panic("Support me!")
	case *ssa.Select:
		panic("Support me!")
	case *ssa.Store:
		panic("Support me!")
	case *ssa.If:
		cond := interpreter.resolveExpression(i.Cond)
		fmt.Printf("[Cond]: %s\n", cond.String())
		out := make([]Interpreter, 0)
		if len(i.Block().Succs) >= 1 {
			fmt.Printf("[Found then branch]\n")
			thenState := *interpreter // копия
			thenState.PathCondition = symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{thenState.PathCondition, cond},
				symbolic.AND,
			)
			fmt.Printf("[thenState PC]: %s\n", thenState.PathCondition.String())
			out = append(out, thenState)
		}
		if len(i.Block().Succs) >= 2 {
			fmt.Printf("[Found else branch]\n")
			elseState := *interpreter // копия
			notCond := symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{cond},
				symbolic.NOT,
			)
			elseState.PathCondition = symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{elseState.PathCondition, notCond},
				symbolic.AND,
			)
			fmt.Printf("[elseState PC]: %s\n", elseState.PathCondition.String())
			out = append(out, elseState)
		}
		return out
	case *ssa.BinOp:
		symExpr := interpreter.resolveExpression(i)
		fmt.Printf("SymExpr: %s\n", symExpr.String())
		return []Interpreter{*interpreter}
	default:
		panic(fmt.Sprintf("Support %T!", i))
	}
}

func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {
	if value == nil {
		fmt.Println("resolveExpression with nil!")
		return nil
	}

	switch v := value.(type) {
	// A Parameter represents an input parameter of a function.
	case *ssa.Parameter:
		// TODO Добавить имя для Ref?? (Issue #9)
		return interpreter.Heap.ParamRef(goTypeToSym(v.Type()))

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
	// TODO IMPL оставшихся?
	/*
	 / One of:
	 / ADD SUB MUL QUO REM          + - * / %
	 / AND OR XOR SHL SHR AND_NOT   & | ^ << >> &^
	 / EQL NEQ LSS LEQ GTR GEQ      == != < <= < >=
	*/
	case *ssa.BinOp:
		l := interpreter.resolveExpression(v.X)
		r := interpreter.resolveExpression(v.Y)

		sop, ok := binOpToSym(v.Op)
		if !ok {
			return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))
		}
		return symbolic.NewBinaryOperation(l, r, sop)

	// Op X
	// One of: NOT SUB ARROW MUL XOR ! - <- * ^
	case *ssa.UnOp:
		// switch x.Op {
		// case token.NOT:
		// return symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{operand}, symbolic.NOT)
		// case token.SUB:
		// у тебя нет унарного минуса как отдельного узла — можно кодировать как (0 - operand)
		// zero := symbolic.NewIntConstant(0)
		// return symbolic.NewBinaryOperation(zero, operand, symbolic.SUB)
		// default:
		// return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))
		// }
		panic("Support me!")

	// The Call instruction represents a function or method call.
	case *ssa.Call:
		// функция + аргументы
		// args := make([]symbolic.SymbolicExpression, 0, len(x.Call.Args))
		// for _, a := range x.Call.Args {
		// args = append(args, valueToSym(a))
		// }
		// name := "call"
		// if x.Call.StaticCallee() != nil {
		// name = x.Call.StaticCallee().Name()
		// }
		// return symbolic.NewFunctionCall(name, args, goTypeToSym(x.Type()))
		panic("Support me!")

	case *ssa.Phi:
		// phi корректнее обрабатывать на входе в блок (по pred), но чтоб “не падало”:
		// return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))
		panic("Support me!")

	default:
		// fallback: просто символ по имени SSA value
		panic("Support me!")
	}
}

func binOpToSym(op token.Token) (symbolic.BinaryOperator, bool) {
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
		return symbolic.MOD, true

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
		return 0, false
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
