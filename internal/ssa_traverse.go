package internal

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"

	"golang.org/x/tools/go/ssa"
)

func traverse_dfs(cur *ssa.BasicBlock, depth int, prevState Interpreter) []Interpreter {
	// печать инструкций блока
	for i, ins := range cur.Instrs {
		fmt.Printf("[%d] %T: %s\n", depth+i, ins, ins.String())
	}

	// если блок пустой — просто возвращаем состояние как “результат”
	if len(cur.Instrs) == 0 {
		return []Interpreter{prevState}
	}

	term := cur.Instrs[len(cur.Instrs)-1]
	out := make([]Interpreter, 0)

	switch t := term.(type) {
	case *ssa.Return, *ssa.Panic:
		fmt.Printf("[exit]: %s\n", term.String())
		out = append(out, prevState)
		return out

	case *ssa.If:
		fmt.Printf("[If]: %s\n", t.Cond.String())

		cond := valueToSym(t.Cond)

		// then
		if len(cur.Succs) >= 1 {
			thenB := cur.Succs[0]

			thenState := prevState // копия
			thenState.PathCondition = symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{thenState.PathCondition, cond},
				symbolic.AND,
			)

			out = append(out, traverse_dfs(thenB, depth+len(cur.Instrs), thenState)...)
		}

		// else
		if len(cur.Succs) >= 2 {
			elseB := cur.Succs[1]

			elseState := prevState // копия
			notCond := symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{cond},
				symbolic.NOT,
			)
			elseState.PathCondition = symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{elseState.PathCondition, notCond},
				symbolic.AND,
			)

			out = append(out, traverse_dfs(elseB, depth+len(cur.Instrs), elseState)...)
		}

		return out

	default:
		// безусловные переходы / прочие терминаторы
		for _, dst := range cur.Succs {
			out = append(out, traverse_dfs(dst, depth+len(cur.Instrs), prevState)...)
		}
		return out
	}
}
