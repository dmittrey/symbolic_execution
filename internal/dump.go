package internal

import (
	"fmt"
	"sort"
	"strings"

	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

func (st Interpreter) Dump() string {
	var b strings.Builder

	fmt.Fprintf(&b, "Interpreter{\n")
	fmt.Fprintf(&b, "  PC: %s\n", safeExpr(st.PathCondition))
	fmt.Fprintf(&b, "  ReturnResults:\n")
	for idx, val := range st.ReturnResults {
		fmt.Fprintf(&b, "    [%d] %s\n", idx, val.String())
	}
	fmt.Fprintf(&b, "  CallStack: %d frame(s)\n", len(st.CallStack))

	for i, fr := range st.CallStack {
		fnName := "<nil>"
		if fr.Function != nil {
			fnName = fr.Function.String()
		}
		fmt.Fprintf(&b, "    Frame[%d]: %s\n", i, fnName)

		// locals: stable order
		if fr.LocalMemory == nil {
			fmt.Fprintf(&b, "      Locals: <nil>\n")
		} else {
			keys := make([]string, 0, len(fr.LocalMemory))
			for k := range fr.LocalMemory {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			fmt.Fprintf(&b, "      Locals: %d\n", len(keys))
			for _, k := range keys {
				fmt.Fprintf(&b, "        %s = %s\n", k, safeExpr(fr.LocalMemory[k]))
			}
		}

		if fr.ReturnValue != nil {
			fmt.Fprintf(&b, "      Return: %s\n", safeExpr(fr.ReturnValue))
		}
	}

	switch mem := st.Heap.(type) {
	case *memory.SymbolicMemory:
		fmt.Fprintf(&b, "  Heap: SymbolicMemory{\n")
		fmt.Fprintf(&b, "    LastIdx: %d\n", mem.LastIdx)
		fmt.Fprintf(&b, "    Params: %d, Allocations: %d\n", len(mem.Params), len(mem.Allocations))

		// Params list
		if len(mem.Params) > 0 {
			fmt.Fprintf(&b, "    ParamsList:\n")
			for i, r := range mem.Params {
				fmt.Fprintf(&b, "      P[%d]: %s (refType=%s)\n", i, r.String(), r.RefType.String())
			}
		}

		// Allocations list
		if len(mem.Allocations) > 0 {
			fmt.Fprintf(&b, "    AllocationsList:\n")
			for i, r := range mem.Allocations {
				fmt.Fprintf(&b, "      A[%d]: %s (refType=%s)\n", i, r.String(), r.RefType.String())
			}
		}

		// ParamNameToRef map (stable order)
		if len(mem.ParamNameToRef) > 0 {
			names := make([]string, 0, len(mem.ParamNameToRef))
			for k := range mem.ParamNameToRef {
				names = append(names, k)
			}
			sort.Strings(names)

			fmt.Fprintf(&b, "    ParamNameToRef:\n")
			for _, name := range names {
				r := mem.ParamNameToRef[name]
				fmt.Fprintf(&b, "      %s -> %s (refType=%s)\n", name, r.String(), r.RefType.String())
			}
		}

		// Segments with stores
		fmt.Fprintf(&b, "    IntSegment: %s\n", dumpArray(mem.IntSegment, 20))
		fmt.Fprintf(&b, "    BoolSegment: %s\n", dumpArray(mem.BoolSegment, 20))

		// ArraysSegment (stable order by key)
		if len(mem.ArraysSegment) == 0 {
			fmt.Fprintf(&b, "    ArraysSegment: <empty>\n")
		} else {
			keys := make([]int, 0, len(mem.ArraysSegment))
			for k := range mem.ArraysSegment {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			fmt.Fprintf(&b, "    ArraysSegment: %d key(s)\n", len(keys))
			for _, k := range keys {
				arr := mem.ArraysSegment[k]
				fmt.Fprintf(&b, "      key=%d: %s\n", k, dumpArray(arr, 20))
			}
		}

		// Structures (если начнёшь их использовать)
		if len(mem.TypeToInt) > 0 || len(mem.RefToTypeIdx) > 0 {
			fmt.Fprintf(&b, "    StructTypes: %d, RefToTypeIdx: %d\n", len(mem.TypeToInt), len(mem.RefToTypeIdx))
		}

		fmt.Fprintf(&b, "  }\n")

		// Heap constraints
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

func dumpArray(a *symbolic.Array, limit int) string {
	if a == nil {
		return "<nil>"
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s (elem=%s, stores=%d)", a.Name, a.ValType.String(), len(a.Stores))

	if len(a.Stores) == 0 {
		return b.String()
	}

	// печатаем последние `limit` store-ов (удобно для циклов)
	start := 0
	if limit > 0 && len(a.Stores) > limit {
		start = len(a.Stores) - limit
		fmt.Fprintf(&b, " [showing last %d]", limit)
	}

	for i := start; i < len(a.Stores); i++ {
		st := a.Stores[i]
		fmt.Fprintf(&b, "\n        store[%d]: idx=%s, val=%s", i, safeExpr(st.Idx), safeExpr(st.Val))
	}

	return b.String()
}

func safeExpr(e symbolic.SymbolicExpression) string {
	if e == nil {
		return "<nil>"
	}
	return e.String()
}
