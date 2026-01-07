package internal

import (
	"fmt"

	"golang.org/x/tools/go/ssa"
)

func traverse_dfs(cur *ssa.BasicBlock, depth int, state Interpreter) {
	for i, ins := range cur.Instrs {
		fmt.Printf("[%d] %T: %s\n", depth+i, ins, ins.String())
	}

	term := cur.Instrs[len(cur.Instrs)-1]

	switch t := term.(type) {
	case *ssa.Return, *ssa.Panic:
		fmt.Printf("[exit]: %s\n", term.String())
		return

	case *ssa.If:
		// t.Cond — условие ветвления (ssa.Value)
		fmt.Printf("[If]: %s\n", t.Cond.String())
		// ADD CONSTRAINT for THEN:  cond
		// ADD CONSTRAINT for ELSE: !cond

		// В go/ssa обычно: Succs[0] = then, Succs[1] = else
		if len(cur.Succs) >= 1 {
			fmt.Printf("[Then]:\n")
			thenB := cur.Succs[0]
			traverse_dfs(thenB, len(cur.Instrs), state)
		}
		if len(cur.Succs) >= 2 {
			fmt.Printf("[Else]:\n")
			elseB := cur.Succs[1]
			traverse_dfs(elseB, len(cur.Instrs), state)
		}
		return

	// case *ssa.Switch:
	// 	// t.Tag — значение, по которому переключаемся (ssa.Value)
	// 	// t.Cases — значения case (обычно константы)
	// 	// ADD CONSTRAINT for CASE i: tag == cases[i]
	// 	// ADD CONSTRAINT for DEFAULT: tag != cases[0] && ... && tag != cases[n-1]

	// 	// В разных версиях детали могут отличаться, но часто Succs соответствует cases + default.
	// 	// Ниже — максимально “тупо”: идём по всем Succs, пытаясь пометить их как case/default.
	// 	nCases := len(t.Cases)
	// 	for i, dst := range cur.Succs {
	// 		if i < nCases {
	// 			step := &PathStep{
	// 				Block:     cur,
	// 				To:        dst,
	// 				Kind:      EdgeSwitchCase,
	// 				CaseValue: t.Cases[i],
	// 			}
	// 			dfs(dst, append(steps, step), depth+1, visits)
	// 		} else {
	// 			step := &PathStep{Block: cur, To: dst, Kind: EdgeSwitchDefault}
	// 			dfs(dst, append(steps, step), depth+1, visits)
	// 		}
	// 	}
	// 	return

	default:
		for _, dst := range cur.Succs {
			traverse_dfs(dst, len(cur.Instrs), state)
		}
		return
	}
}
