package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"

	"symbolic-execution-course/internal/symbolic"

	"golang.org/x/tools/go/ssa"
)

// --- helpers: types ---

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

// --- main: ssa.Value -> SymbolicExpression ---

func valueToSym(v ssa.Value) symbolic.SymbolicExpression {
	if v == nil {
		return nil
	}

	switch x := v.(type) {
	case *ssa.Const:
		// nil, true/false, int, string
		if x.Value == nil {
			// как обработать nil — зависит от модели памяти/refs.
			// Пока сделаем символ.
			return symbolic.NewSymbolicVariable("nil", symbolic.RefType)
		}

		switch x.Value.Kind() {
		case constant.Bool:
			return symbolic.NewBoolConstant(constant.BoolVal(x.Value))

		case constant.Int:
			i, ok := constant.Int64Val(x.Value)
			if !ok {
				panic(fmt.Sprintf("big int const not supported: %s", x.Value.String()))
			}
			return symbolic.NewIntConstant(i)

		case constant.String:
			return symbolic.NewStringConstant(constant.StringVal(x.Value))

		default:
			// float/complex etc — если не надо, можно паникнуть
			return symbolic.NewSymbolicVariable(x.String(), goTypeToSym(x.Type()))
		}

	case *ssa.Parameter:
		return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))

	case *ssa.Global:
		return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))

	case *ssa.BinOp:
		l := valueToSym(x.X)
		r := valueToSym(x.Y)

		sop, ok := binOpToSym(x.Op)
		if !ok {
			// например &&/|| в SSA часто идут как If на cond, а не как BinOp,
			// но если встретишь — можно расширить.
			return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))
		}
		return symbolic.NewBinaryOperation(l, r, sop)

	case *ssa.UnOp:
		// в SSA: UnOp может быть ! (NOT), - (SUB в токенах как unary), * (load), <- (recv)
		operand := valueToSym(x.X)

		switch x.Op {
		case token.NOT:
			return symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{operand}, symbolic.NOT)
		case token.SUB:
			// у тебя нет унарного минуса как отдельного узла — можно кодировать как (0 - operand)
			zero := symbolic.NewIntConstant(0)
			return symbolic.NewBinaryOperation(zero, operand, symbolic.SUB)
		default:
			return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))
		}

	case *ssa.Call:
		// функция + аргументы
		args := make([]symbolic.SymbolicExpression, 0, len(x.Call.Args))
		for _, a := range x.Call.Args {
			args = append(args, valueToSym(a))
		}
		name := "call"
		if x.Call.StaticCallee() != nil {
			name = x.Call.StaticCallee().Name()
		}
		return symbolic.NewFunctionCall(name, args, goTypeToSym(x.Type()))

	case *ssa.Phi:
		// phi корректнее обрабатывать на входе в блок (по pred), но чтоб “не падало”:
		return symbolic.NewSymbolicVariable(x.Name(), goTypeToSym(x.Type()))

	default:
		// fallback: просто символ по имени SSA value
		// (для Register-like значений Name() часто ок)
		return symbolic.NewSymbolicVariable(v.Name(), goTypeToSym(v.Type()))
	}
}
