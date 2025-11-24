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
	// 1) UNSAT: int prevent aliasing
	runTest("1) UNSAT: int prevent aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.Allocate(symbolic.IntType)
		foo2 := mem.Allocate(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(foo2, symbolic.NewIntConstant(5))
		mem.AssignValue(foo1, symbolic.NewIntConstant(2))

		// Получаем символьное выражение "значение по foo2"
		valFoo2 := mem.GetValue(foo2)

		// Строим условие valFoo2 == 5 (SAT)
		cond := symbolic.NewBinaryOperation(
			valFoo2,
			symbolic.NewIntConstant(2),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		// fmt.Println("== Z3 formula dump (int) ==")
		// fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 2) SAT: int assign
	runTest("2) SAT: int assign", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.Allocate(symbolic.IntType)
		foo2 := mem.Allocate(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(foo2, symbolic.NewIntConstant(5))
		mem.AssignValue(foo1, symbolic.NewIntConstant(2))

		// Получаем символьное выражение "значение по foo2"
		valFoo2 := mem.GetValue(foo2)

		// Строим условие valFoo2 == 5 (SAT)
		cond := symbolic.NewBinaryOperation(
			valFoo2,
			symbolic.NewIntConstant(5),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		// fmt.Println("== Z3 formula dump (int) ==")
		// fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 3) UNSAT: bool prevent aliasing
	runTest("3) UNSAT: bool prevent aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		b1 := mem.Allocate(symbolic.BoolType)
		b2 := mem.Allocate(symbolic.BoolType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(b2, symbolic.NewBoolConstant(true))
		mem.AssignValue(b1, symbolic.NewBoolConstant(false))

		// Получаем символьное выражение "значение по b2"
		valB2 := mem.GetValue(b2)

		// Строим условие valB2 == true (SAT)
		cond := symbolic.NewBinaryOperation(
			valB2,
			symbolic.NewBoolConstant(false),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		// fmt.Println("== Z3 formula dump (bool) ==")
		// fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 4) SAT: bool assign
	runTest("4) SAT: bool assign", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		b1 := mem.Allocate(symbolic.BoolType)
		b2 := mem.Allocate(symbolic.BoolType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(b2, symbolic.NewBoolConstant(true))
		mem.AssignValue(b1, symbolic.NewBoolConstant(false))

		// Получаем символьное выражение "значение по b2"
		valB2 := mem.GetValue(b2)

		// Строим условие valB2 == true (SAT)
		cond := symbolic.NewBinaryOperation(
			valB2,
			symbolic.NewBoolConstant(true),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)

		// fmt.Println("== Z3 formula dump (bool) ==")
		// fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})
}
