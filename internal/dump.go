package internal

import (
	"fmt"
	"strings"

	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

func (st Interpreter) Dump() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Interpreter{\n")
	fmt.Fprintf(&b, "  PC: %s\n", safeExpr(st.PathCondition))
	fmt.Fprintf(&b, "  CallStack: %d frame(s)\n", len(st.CallStack))

	for i, fr := range st.CallStack {
		fnName := "<nil>"
		if fr.Function != nil {
			fnName = fr.Function.String()
		}
		fmt.Fprintf(&b, "    Frame[%d]: %s\n", i, fnName)

		fmt.Fprintf(&b, "      Locals: %d\n", len(fr.LocalMemory))
		for k, v := range fr.LocalMemory {
			fmt.Fprintf(&b, "        %s = %s\n", k, safeExpr(v))
		}

		if fr.ReturnValue != nil {
			fmt.Fprintf(&b, "      Return: %s\n", safeExpr(fr.ReturnValue))
		}
	}

	// Если Heap = *memory.SymbolicMemory — можно вытащить его ограничения и сегменты.
	switch mem := st.Heap.(type) {
	case *memory.SymbolicMemory:
		fmt.Fprintf(&b, "  Heap: SymbolicMemory{\n")
		fmt.Fprintf(&b, "    Params: %d, Allocations: %d\n", len(mem.Params), len(mem.Allocations))
		fmt.Fprintf(&b, "    IntSegment: %s\n", mem.IntSegment.String())
		fmt.Fprintf(&b, "    BoolSegment: %s\n", mem.BoolSegment.String())
		fmt.Fprintf(&b, "    ArraysSegment keys: %d\n", len(mem.ArraysSegment))
		fmt.Fprintf(&b, "  }\n")

		// constraints (если хочешь)
		cs := mem.Constraints()
		fmt.Fprintf(&b, "  HeapConstraints: %d\n", len(cs))
		if len(cs) > 0 {
			for i, c := range cs {
				fmt.Fprintf(&b, "    C[%d]: %s\n", i, safeExpr(c))
			}
		}

	default:
		fmt.Fprintf(&b, "  Heap: %T\n", st.Heap)
	}

	fmt.Fprintf(&b, "}\n")
	return b.String()
}

func safeExpr(e symbolic.SymbolicExpression) string {
	if e == nil {
		return "<nil>"
	}
	return e.String()
}
