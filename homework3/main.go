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
	// 1) UNSAT: int prevent aliasing
	runTest("1) UNSAT: int prevent aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.AllocatePrimitiveArray(symbolic.IntType)
		foo2 := mem.AllocatePrimitiveArray(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignToArray(foo2, symbolic.NewIntConstant(1), symbolic.NewIntConstant(5))
		fmt.Println(mem.IntSegment.String())
		mem.AssignToArray(foo1, symbolic.NewIntConstant(1), symbolic.NewIntConstant(2))
		fmt.Println(mem.IntSegment.String())

		// Строим условие valFoo2 == 5 (SAT)
		cond := symbolic.NewBinaryOperation(
			mem.GetFromArray(foo2, symbolic.NewIntConstant(1)),
			symbolic.NewIntConstant(2),
			symbolic.EQ)

		tr.AssertExpr(cond)
		for _, v := range mem.Constraints() {
			tr.AssertExpr(v)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			panic("")
		}
	})

	// 2) SAT: int aliasing
	runTest("2) SAT: int aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.ParamPrimitiveArray(symbolic.IntType)
		foo2 := mem.ParamPrimitiveArray(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignToArray(foo2, symbolic.NewIntConstant(1), symbolic.NewIntConstant(5))
		fmt.Println(mem.IntSegment.String())
		mem.AssignToArray(foo1, symbolic.NewIntConstant(1), symbolic.NewIntConstant(2))
		fmt.Println(mem.IntSegment.String())

		// Строим условие valFoo2 == 5 (SAT)
		cond := symbolic.NewBinaryOperation(
			mem.GetFromArray(foo2, symbolic.NewIntConstant(1)),
			symbolic.NewIntConstant(2),
			symbolic.EQ)

		tr.AssertExpr(cond)
		for _, v := range mem.Constraints() {
			tr.AssertExpr(v)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m == nil {
			panic("")
		}
	})

	// 3) SAT: bool[][] aliasing
	runTest("3) SAT: bool[][] aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// bool[][]
		foo1 := mem.AllocateComplexArray(1, symbolic.BoolType)
		foo2 := mem.AllocateComplexArray(1, symbolic.BoolType)

		// Присваиваем конкретные символические значения
		// foo2[1][1] = true
		foo2_1 := mem.GetFromArray(foo2, symbolic.NewIntConstant(1))
		if foo2_1_r, ok := foo2_1.(*symbolic.Ref); ok {
			mem.AssignToArray(foo2_1_r, symbolic.NewIntConstant(1), symbolic.NewBoolConstant(true))
		}
		fmt.Println(mem.IntSegment.String())

		// foo1[1][1] = false
		foo1_1 := mem.GetFromArray(foo1, symbolic.NewIntConstant(1))
		if foo1_1_r, ok := foo1_1.(*symbolic.Ref); ok {
			mem.AssignToArray(foo1_1_r, symbolic.NewIntConstant(1), symbolic.NewBoolConstant(false))
		}
		fmt.Println(mem.IntSegment.String())

		// Строим условие foo2[1][1] == false (SAT)
		foo2_1 = mem.GetFromArray(foo2, symbolic.NewIntConstant(1))
		if foo2_1_r, ok := foo2_1.(*symbolic.Ref); ok {
			cond := symbolic.NewBinaryOperation(
				mem.GetFromArray(foo2_1_r, symbolic.NewIntConstant(1)),
				symbolic.NewBoolConstant(false),
				symbolic.EQ)

			tr.AssertExpr(cond)
		}
		for _, v := range mem.Constraints() {
			tr.AssertExpr(v)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m == nil {
			panic("")
		}
	})

	// 4) UNSAT: int prevent aliasing
	runTest("4) UNSAT: int prevent aliasing", func(tr *translator.Z3Translator) {
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
		for _, c := range mem.Constraints() {
			tr.AssertExpr(c)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			panic("")
		}
	})

	// 5) SAT: int aliasing
	runTest("5) SAT: int aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		foo1 := mem.Param(symbolic.IntType)
		foo2 := mem.Param(symbolic.IntType)

		// Присваиваем конкретные символические значения
		mem.AssignValue(foo2, symbolic.NewIntConstant(5))
		mem.AssignValue(foo1, symbolic.NewIntConstant(2))

		// Получаем символьное выражение "значение по foo2"
		valFoo2 := mem.GetValue(foo2)

		// Строим условие valFoo2 == 2 (SAT)
		cond := symbolic.NewBinaryOperation(
			valFoo2,
			symbolic.NewIntConstant(2),
			symbolic.EQ,
		)

		tr.AssertExpr(cond)
		for _, c := range mem.Constraints() {
			tr.AssertExpr(c)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		} else {
			panic("")
		}
	})

	// 6) SAT: Foo[]
	// struct Foo {
	// 	idx int
	// }
	runTest("6) SAT: Foo[]", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// Foo[] foo1 = []
		// Foo[] foo2 = []
		foo1 := mem.AllocateStructsArray(0, "Foo", []symbolic.ExpressionType{symbolic.BoolType})
		foo2 := mem.AllocateStructsArray(0, "Foo", []symbolic.ExpressionType{symbolic.BoolType})

		// foo1_1 = foo1[1]
		// foo2_1 = foo2[1]
		foo1_1 := mem.GetFromArray(foo1, symbolic.NewIntConstant(1))
		foo2_1 := mem.GetFromArray(foo2, symbolic.NewIntConstant(1))

		// foo1_1.idx = true
		if foo1_1_r, ok := foo1_1.(*symbolic.Ref); ok {
			foo1_1_r0 := mem.FieldAddress(foo1_1_r, 0)
			mem.AssignValue(foo1_1_r0, symbolic.NewBoolConstant(true))
		}
		// foo2_1.idx = false
		if foo2_1_r, ok := foo2_1.(*symbolic.Ref); ok {
			foo2_1_r0 := mem.FieldAddress(foo2_1_r, 0)
			mem.AssignValue(foo2_1_r0, symbolic.NewBoolConstant(false))
		}

		// Строим условие foo1_1.idx == false (SAT)
		if foo1_1_r, ok := foo1_1.(*symbolic.Ref); ok {
			foo1_1_r0 := mem.FieldAddress(foo1_1_r, 0)
			val := mem.GetValue(foo1_1_r0)
			cond := symbolic.NewBinaryOperation(
				val,
				symbolic.NewBoolConstant(false),
				symbolic.EQ,
			)
			tr.AssertExpr(cond)
		}

		// tr.AssertExpr(cond)
		for _, c := range mem.Constraints() {
			tr.AssertExpr(c)
		}

		fmt.Println("== Z3 formula dump (int) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		} else {
			// panic("")
		}
	})

	runTest("ParamStruct: SAT by aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// Two struct params can alias
		s1 := mem.ParamStruct("Foo", []symbolic.ExpressionType{symbolic.BoolType})
		s2 := mem.ParamStruct("Foo", []symbolic.ExpressionType{symbolic.BoolType})

		// s1.f0 = true
		s1f0 := mem.FieldAddress(s1, 0)
		mem.AssignValue(s1f0, symbolic.NewBoolConstant(true))

		// s2.f0 = false
		s2f0 := mem.FieldAddress(s2, 0)
		mem.AssignValue(s2f0, symbolic.NewBoolConstant(false))

		// Assert: s1.f0 == false  (SAT if s1 and s2 alias)
		val := mem.GetValue(s1f0)
		cond := symbolic.NewBinaryOperation(val, symbolic.NewBoolConstant(false), symbolic.EQ)
		tr.AssertExpr(cond)

		for _, c := range mem.Constraints() {
			tr.AssertExpr(c)
		}

		fmt.Println(tr.SolverToString())
		if m := mustSat(tr); m == nil {
			panic("expected SAT")
		}
	})

	runTest("ParamStructsArray: SAT by object aliasing under distinct bases", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		a1 := mem.ParamStructsArray(0, "Foo", []symbolic.ExpressionType{symbolic.BoolType})
		a2 := mem.ParamStructsArray(0, "Foo", []symbolic.ExpressionType{symbolic.BoolType})

		// Force array base pointers to be different (but allow object aliasing via ArraySegment)
		tr.AssertExpr(symbolic.NewBinaryOperation(a1.Idx, a2.Idx, symbolic.NE))

		// a1_1 := a1[1]
		a1_1 := mem.GetFromArray(a1, symbolic.NewIntConstant(1)).(*symbolic.Ref)
		a1_1_f0 := mem.FieldAddress(a1_1, 0)
		mem.AssignValue(a1_1_f0, symbolic.NewBoolConstant(true))

		// a2_2 := a2[2]
		a2_2 := mem.GetFromArray(a2, symbolic.NewIntConstant(2)).(*symbolic.Ref)
		a2_2_f0 := mem.FieldAddress(a2_2, 0)
		mem.AssignValue(a2_2_f0, symbolic.NewBoolConstant(false))

		// Assert: a1[1].f0 == false (SAT if they alias to same object handle)
		val := mem.GetValue(a1_1_f0)
		cond := symbolic.NewBinaryOperation(val, symbolic.NewBoolConstant(false), symbolic.EQ)
		tr.AssertExpr(cond)

		for _, c := range mem.Constraints() {
			tr.AssertExpr(c)
		}

		fmt.Println(tr.SolverToString())
		if m := mustSat(tr); m == nil {
			panic("expected SAT")
		}
	})

}
