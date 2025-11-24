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

/*
func OnStackInt() *int:
0:                                                                entry P:0 S:0
        t0 = new int (foo)                                                 *int
        *t0 = 5:int
        return t0

func OnHeapInt() *int:
0:                                                                entry P:0 S:0
        t0 = new int (new)                                                 *int
        *t0 = 5:int
        return t0

*/

func main() {
	// 1) SAT: значение по ссылке foo2 равно 5
	runTest("1) SAT (int, GetValue == 5)", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.Allocate(symbolic.IntType)
		foo2 := mem.Allocate(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(foo2, symbolic.NewIntConstant(5))
		mem.AssignValue(foo1, symbolic.NewIntConstant(2))

		// Получаем символьное выражение "значение по foo2"
		valFoo2 := mem.GetValue(foo2)

		// Строим условие valFoo2 == 2
		cond := symbolic.NewBinaryOperation(
			valFoo2,
			symbolic.NewIntConstant(2),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		fmt.Println("== Z3 formula dump ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})
}
