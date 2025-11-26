package main

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"

	"github.com/ebukreev/go-z3/z3"
)

func runTest(name string, test func(tr *translator.Z3Translator)) {
	fmt.Printf("\n== %s ==\n", name)
	tr := translator.NewZ3Translator()
	defer tr.Close()
	test(tr)
}

func mustSat(tr *translator.Z3Translator) *z3.Model {
	sat, err := tr.IsSatisfiable()
	if err != nil {
		log.Fatalf("Check error: %v", err)
	}
	if !sat {
		fmt.Println("UNSAT")
		return nil
	}
	fmt.Println("SAT")
	return tr.Model()
}

func main() {
	// 12) Sat
	runTest("12) SAT", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// Аналогично двум разным указателям foo1, foo2
		foo1 := mem.ParamRef(symbolic.BoolType)
		foo2 := mem.ParamRef(symbolic.BoolType)

		// foo2.a = 5
		mem.AssignField(foo2, 0, symbolic.NewIntConstant(5))
		// foo1.a = 2
		mem.AssignField(foo1, 0, symbolic.NewIntConstant(2))

		// if foo2.a == 2 { ... }
		foo2a := mem.GetFieldValue(foo2, 0)

		cond := symbolic.NewBinaryOperation(
			foo2a,
			symbolic.NewIntConstant(2),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		fmt.Println("== Z3 formula dump (Foo no-alias) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		} else {
			panic("hello")
		}
	})

}
