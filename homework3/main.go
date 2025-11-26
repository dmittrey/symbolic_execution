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

	// 5) SAT: struct Person basic fields
	runTest("5) SAT: struct Person basic fields", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// t0 = local Person (p) *Person
		person := mem.AllocateStruct("Person", []symbolic.ExpressionType{
			symbolic.StringType, // Name
			symbolic.IntType,    // Age
			symbolic.IntType,    // ID
		})

		// t1 = &t0.Name [#0]; *t1 = "Alice"
		mem.AssignField(person, 0, symbolic.NewStringConstant("Alice"))

		// t2 = &t0.Age [#1]; *t2 = 25
		mem.AssignField(person, 1, symbolic.NewIntConstant(25))

		// t3 = &t0.ID [#2]; *t3 = 1001
		mem.AssignField(person, 2, symbolic.NewIntConstant(1001))

		// t4 = *t0  (по сути мы просто читаем поля у person)
		nameVal := mem.GetFieldValue(person, 0)
		ageVal := mem.GetFieldValue(person, 1)
		idVal := mem.GetFieldValue(person, 2)

		// Проверяем:
		// Name == "Alice"
		tr.AssertExpr(symbolic.NewBinaryOperation(
			nameVal,
			symbolic.NewStringConstant("Alice"),
			symbolic.EQ,
		))

		// Age == 25
		tr.AssertExpr(symbolic.NewBinaryOperation(
			ageVal,
			symbolic.NewIntConstant(25),
			symbolic.EQ,
		))

		// ID == 1001
		tr.AssertExpr(symbolic.NewBinaryOperation(
			idVal,
			symbolic.NewIntConstant(1001),
			symbolic.EQ,
		))

		fmt.Println("== Z3 formula dump (bool) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 6) SAT: struct Person modification (p.Age = p.Age + 1, p.ID = p.ID * 2)
	runTest("6) SAT: struct Person modification", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// t0 = local Person (p) *Person
		p := mem.AllocateStruct("Person", []symbolic.ExpressionType{
			symbolic.StringType,
			symbolic.IntType,
			symbolic.IntType,
		})

		// Параметр p Person: поля как символьные переменные
		name0 := symbolic.NewSymbolicVariable("p_name0", symbolic.StringType)
		age0 := symbolic.NewSymbolicVariable("p_age0", symbolic.IntType)
		id0 := symbolic.NewSymbolicVariable("p_id0", symbolic.IntType)

		mem.AssignField(p, 0, name0)
		mem.AssignField(p, 1, age0)
		mem.AssignField(p, 2, id0)

		// t1 = &t0.Age; t2 = *t1; t3 = t2 + 1; *Age = t3
		origAge := mem.GetFieldValue(p, 1)
		agePlus1 := symbolic.NewBinaryOperation(
			origAge,
			symbolic.NewIntConstant(1),
			symbolic.ADD,
		)
		mem.AssignField(p, 1, agePlus1)

		// t5 = &t0.ID; t6 = *t5; t7 = t6 * 2; *ID = t7
		origID := mem.GetFieldValue(p, 2)
		idTimes2 := symbolic.NewBinaryOperation(
			origID,
			symbolic.NewIntConstant(2),
			symbolic.MUL,
		)
		mem.AssignField(p, 2, idTimes2)

		// t9 = *t0; return t9
		finalName := mem.GetFieldValue(p, 0)
		finalAge := mem.GetFieldValue(p, 1)
		finalID := mem.GetFieldValue(p, 2)

		// Проверяем семантику:
		// finalName == name0
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalName,
			name0,
			symbolic.EQ,
		))
		// finalAge == age0 + 1
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalAge,
			symbolic.NewBinaryOperation(age0, symbolic.NewIntConstant(1), symbolic.ADD),
			symbolic.EQ,
		))
		// finalID == id0 * 2
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalID,
			symbolic.NewBinaryOperation(id0, symbolic.NewIntConstant(2), symbolic.MUL),
			symbolic.EQ,
		))

		fmt.Println("== Z3 formula dump (struct modification) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 7) SAT: testStructPointer
	runTest("7) SAT: testStructPointer", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// p := &Person{Name: "Bob", Age: 30, ID: 2002}
		p := mem.AllocateStruct("Person", []symbolic.ExpressionType{
			symbolic.StringType,
			symbolic.IntType,
			symbolic.IntType,
		})

		mem.AssignField(p, 0, symbolic.NewStringConstant("Bob"))
		mem.AssignField(p, 1, symbolic.NewIntConstant(30))
		mem.AssignField(p, 2, symbolic.NewIntConstant(2002))

		// p.Age = p.Age + 5
		oldAge := mem.GetFieldValue(p, 1)
		agePlus5 := symbolic.NewBinaryOperation(
			oldAge,
			symbolic.NewIntConstant(5),
			symbolic.ADD,
		)
		mem.AssignField(p, 1, agePlus5)

		// return p
		finalName := mem.GetFieldValue(p, 0)
		finalAge := mem.GetFieldValue(p, 1)
		finalID := mem.GetFieldValue(p, 2)

		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalName,
			symbolic.NewStringConstant("Bob"),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalAge,
			symbolic.NewIntConstant(35),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalID,
			symbolic.NewIntConstant(2002),
			symbolic.EQ,
		))

		fmt.Println("== Z3 formula dump (testStructPointer) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 8) SAT: testStructPointerModification
	runTest("8) SAT: testStructPointerModification", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// параметр p *Person, моделируем как один Person в памяти
		p := mem.AllocateStruct("Person", []symbolic.ExpressionType{
			symbolic.StringType,
			symbolic.IntType,
			symbolic.IntType,
		})

		name0 := symbolic.NewSymbolicVariable("pp_name0", symbolic.StringType)
		age0 := symbolic.NewSymbolicVariable("pp_age0", symbolic.IntType)
		id0 := symbolic.NewSymbolicVariable("pp_id0", symbolic.IntType)

		mem.AssignField(p, 0, name0)
		mem.AssignField(p, 1, age0)
		mem.AssignField(p, 2, id0)

		// if p != nil { ... } — берём ветку p != nil

		// p.Age = p.Age + 10
		oldAge := mem.GetFieldValue(p, 1)
		agePlus10 := symbolic.NewBinaryOperation(
			oldAge,
			symbolic.NewIntConstant(10),
			symbolic.ADD,
		)
		mem.AssignField(p, 1, agePlus10)

		// p.ID = p.ID + 1000
		oldID := mem.GetFieldValue(p, 2)
		idPlus1000 := symbolic.NewBinaryOperation(
			oldID,
			symbolic.NewIntConstant(1000),
			symbolic.ADD,
		)
		mem.AssignField(p, 2, idPlus1000)

		finalAge := mem.GetFieldValue(p, 1)
		finalID := mem.GetFieldValue(p, 2)

		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalAge,
			symbolic.NewBinaryOperation(age0, symbolic.NewIntConstant(10), symbolic.ADD),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalID,
			symbolic.NewBinaryOperation(id0, symbolic.NewIntConstant(1000), symbolic.ADD),
			symbolic.EQ,
		))

		fmt.Println("== Z3 formula dump (testStructPointerModification) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 9) SAT: testNestedStructPointer
	runTest("9) SAT: testNestedStructPointer", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// p := testStructPointer(); testStructPointerModification(p)
		p := mem.AllocateStruct("Person", []symbolic.ExpressionType{
			symbolic.StringType,
			symbolic.IntType,
			symbolic.IntType,
		})

		// testStructPointer: конструктор + p.Age = p.Age + 5
		mem.AssignField(p, 0, symbolic.NewStringConstant("Bob"))
		mem.AssignField(p, 1, symbolic.NewIntConstant(30))
		mem.AssignField(p, 2, symbolic.NewIntConstant(2002))

		oldAge1 := mem.GetFieldValue(p, 1)
		agePlus5 := symbolic.NewBinaryOperation(
			oldAge1,
			symbolic.NewIntConstant(5),
			symbolic.ADD,
		)
		mem.AssignField(p, 1, agePlus5) // Age = 35

		// testStructPointerModification(p): +10 к Age и +1000 к ID
		oldAge2 := mem.GetFieldValue(p, 1)
		agePlus10 := symbolic.NewBinaryOperation(
			oldAge2,
			symbolic.NewIntConstant(10),
			symbolic.ADD,
		)
		mem.AssignField(p, 1, agePlus10) // Age = 45

		oldID := mem.GetFieldValue(p, 2)
		idPlus1000 := symbolic.NewBinaryOperation(
			oldID,
			symbolic.NewIntConstant(1000),
			symbolic.ADD,
		)
		mem.AssignField(p, 2, idPlus1000) // ID = 3002

		finalName := mem.GetFieldValue(p, 0)
		finalAge := mem.GetFieldValue(p, 1)
		finalID := mem.GetFieldValue(p, 2)

		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalName,
			symbolic.NewStringConstant("Bob"),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalAge,
			symbolic.NewIntConstant(45),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			finalID,
			symbolic.NewIntConstant(3002),
			symbolic.EQ,
		))

		fmt.Println("== Z3 formula dump (testNestedStructPointer) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 10) SAT: testArrayFixed — arr[i] = i*i
	runTest("10) SAT: testArrayFixed", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		arr := mem.AllocateArray(symbolic.IntType)

		for i := 0; i < 5; i++ {
			mem.AssignToArray(arr, i, symbolic.NewIntConstant(int64(i*i)))
		}

		// проверяем, что arr[i] == i*i
		for i := 0; i < 5; i++ {
			val := mem.GetFromArray(arr, i)
			tr.AssertExpr(symbolic.NewBinaryOperation(
				val,
				symbolic.NewIntConstant(int64(i*i)),
				symbolic.EQ,
			))
		}

		fmt.Println("== Z3 formula dump (testArrayFixed) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 11) SAT: testArrayModification — arr'[i] = arr[i] + 1
	runTest("11) SAT: testArrayModification", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// параметр arr [5]int: задаём 5 символьных значений
		arr := mem.AllocateArray(symbolic.IntType)

		for i := 0; i < 5; i++ {
			v := symbolic.NewSymbolicVariable(fmt.Sprintf("arr_%d", i), symbolic.IntType)
			mem.AssignToArray(arr, i, v)
		}

		// тело функции: для каждого i arr[i] = arr[i] + 1
		for i := 0; i < 5; i++ {
			oldVal := mem.GetFromArray(arr, i)
			newVal := symbolic.NewBinaryOperation(
				oldVal,
				symbolic.NewIntConstant(1),
				symbolic.ADD,
			)
			mem.AssignToArray(arr, i, newVal)
		}

		// ожидаем: finalArr[i] == arr_i + 1
		for i := 0; i < 5; i++ {
			finalVal := mem.GetFromArray(arr, i)
			expected := symbolic.NewBinaryOperation(
				symbolic.NewSymbolicVariable(fmt.Sprintf("arr_%d", i), symbolic.IntType),
				symbolic.NewIntConstant(1),
				symbolic.ADD,
			)
			tr.AssertExpr(symbolic.NewBinaryOperation(
				finalVal,
				expected,
				symbolic.EQ,
			))
		}

		fmt.Println("== Z3 formula dump (testArrayModification) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

	// 12) UNSAT: struct Foo prevent aliasing
	runTest("12) UNSAT: struct Foo prevent aliasing", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// Аналогично двум разным указателям foo1, foo2
		foo1 := mem.AllocateStruct("Foo", []symbolic.ExpressionType{
			symbolic.IntType, // a
		})
		foo2 := mem.AllocateStruct("Foo", []symbolic.ExpressionType{
			symbolic.IntType, // a
		})

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
		}
	})

	// 13) SAT: struct Foo aliasing (foo1 == foo2)
	runTest("13) SAT: struct Foo aliasing (same ref)", func(tr *translator.Z3Translator) {
		mem := memory.NewSymbolicMemory()

		// Один и тот же объект, два "параметра" указывают на него
		foo1 := mem.AllocateStruct("Foo", []symbolic.ExpressionType{
			symbolic.IntType, // a
		})
		foo2 := foo1 // алиас

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

		fmt.Println("== Z3 formula dump (Foo alias) ==")
		fmt.Println(tr.SolverToString())

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model:\n%s\n", m.String())
		}
	})

}
