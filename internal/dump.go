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
		fmt.Fprintf(&b, "    [%d] %s\n", idx, safeExpr(val))
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
		fmt.Fprintf(&b, "    LastRefIdx: %d\n", mem.LastRefIdx)
		fmt.Fprintf(&b, "    Params: %d, Allocations: %d\n", len(mem.Params), len(mem.Allocations))

		// Params list
		if len(mem.Params) > 0 {
			fmt.Fprintf(&b, "    ParamsList:\n")
			for i, r := range mem.Params {
				fmt.Fprintf(&b, "      P[%d]: %s\n", i, dumpRef(r))
			}
		}

		// Allocations list
		if len(mem.Allocations) > 0 {
			fmt.Fprintf(&b, "    AllocationsList:\n")
			for i, r := range mem.Allocations {
				fmt.Fprintf(&b, "      A[%d]: %s\n", i, dumpRef(r))
			}
		}

		// Core segments
		fmt.Fprintf(&b, "    IntSegment:    %s\n", dumpArray(&mem.IntSegment, 30))
		fmt.Fprintf(&b, "    BoolSegment:   %s\n", dumpArray(&mem.BoolSegment, 30))
		fmt.Fprintf(&b, "    StringSegment: %s\n", dumpArray(&mem.StringSegment, 30))
		fmt.Fprintf(&b, "    ArraySegment:  %s\n", dumpArray(&mem.ArraySegment, 30))

		// Struct segments
		if len(mem.StructSegments) == 0 {
			fmt.Fprintf(&b, "    StructSegments: <empty>\n")
		} else {
			typeNames := make([]string, 0, len(mem.StructSegments))
			for tn := range mem.StructSegments {
				typeNames = append(typeNames, tn)
			}
			sort.Strings(typeNames)

			fmt.Fprintf(&b, "    StructSegments: %d type(s)\n", len(typeNames))
			for _, tn := range typeNames {
				fields := mem.StructSegments[tn]

				// stable field order: 0..n-1 (если дырки — сортируем ключи)
				fieldIdxs := make([]int, 0, len(fields))
				for idx := range fields {
					fieldIdxs = append(fieldIdxs, idx)
				}
				sort.Ints(fieldIdxs)

				fmt.Fprintf(&b, "      %s:\n", tn)
				for _, fi := range fieldIdxs {
					fmt.Fprintf(&b, "        field[%d]: %s\n", fi, dumpArray(fields[fi], 30))
				}
			}
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

func dumpRef(r *symbolic.Ref) string {
	if r == nil {
		return "<nil ref>"
	}

	switch r.RefType {
	case symbolic.ArrayType:
		// ArrayDimension: 0 => "T[]" (примитивный/структурный массив)
		// ArrayDimension: 1 => "T[][]" (и т.д.)
		dim := r.ArrayDimension
		base := r.ArrayBaseType.String()

		// Если массив структур — напечатаем имя типа
		if r.ArrayBaseType == symbolic.StructType && r.StructTypeName != "" {
			base = fmt.Sprintf("struct(%s)", r.StructTypeName)
		}

		return fmt.Sprintf("%s (refType=array, dim=%d, base=%s)", r.String(), dim, base)

	case symbolic.StructType:
		// Поле может быть ещё не выбранным (StructFieldIdx=-1)
		if r.StructTypeName == "" {
			return fmt.Sprintf("%s (refType=struct, type=<unset>, field=%d)", r.String(), r.StructFieldIdx)
		}
		return fmt.Sprintf("%s (refType=struct, type=%s, field=%d)", r.String(), r.StructTypeName, r.StructFieldIdx)

	default:
		return fmt.Sprintf("%s (refType=%s)", r.String(), r.RefType.String())
	}
}

func dumpArray(a *symbolic.Array, limit int) string {
	if a == nil {
		return "<nil>"
	}

	// посчитаем типы операций
	var stores, selects int
	for _, op := range a.Ops {
		switch op.Kind {
		case symbolic.OpStore:
			stores++
		case symbolic.OpSelect:
			selects++
		}
	}

	var b strings.Builder
	fmt.Fprintf(&b, "%s (elem=%s, ops=%d, store=%d, select=%d)", a.Name, a.ElemType.String(), len(a.Ops), stores, selects)

	if len(a.Ops) == 0 {
		return b.String()
	}

	start := 0
	if limit > 0 && len(a.Ops) > limit {
		start = len(a.Ops) - limit
		fmt.Fprintf(&b, " [showing last %d]", limit)
	}

	for i := start; i < len(a.Ops); i++ {
		op := a.Ops[i]
		switch op.Kind {
		case symbolic.OpStore:
			fmt.Fprintf(&b, "\n        op[%d]: store idx=%s, val=%s", i, safeExpr(op.Idx), safeExpr(op.Val))
		case symbolic.OpSelect:
			// op.Val — это "результат чтения" (символьная переменная)
			fmt.Fprintf(&b, "\n        op[%d]: select idx=%s, res=%s", i, safeExpr(op.Idx), safeExpr(op.Val))
		default:
			fmt.Fprintf(&b, "\n        op[%d]: <unknown-op> idx=%s, val=%s", i, safeExpr(op.Idx), safeExpr(op.Val))
		}
	}

	return b.String()
}

func safeExpr(e symbolic.SymbolicExpression) string {
	if e == nil {
		return "<nil>"
	}
	return e.String()
}
