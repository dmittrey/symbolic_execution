package main

import (
	"fmt"
	"log"

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
	fmt.Println("=== Symbolic Expressions Demo (no Unary) ===")

	// 1) Арифметика: x + y > 5
	runTest("1) Arithmetic: x + y > 5", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		y := symbolic.NewSymbolicVariable("y", symbolic.IntType)
		sum := symbolic.NewBinaryOperation(x, y, symbolic.ADD)
		gt5 := symbolic.NewBinaryOperation(sum, symbolic.NewIntConstant(5), symbolic.GT)
		tr.AssertExpr(gt5)

		if m := mustSat(tr); m != nil {
			xv, _ := tr.GetIntValue(m, x)
			yv, _ := tr.GetIntValue(m, y)
			fmt.Printf("Model: x=%d, y=%d\n", xv, yv)
		}
	})

	// 2) Логика AND: (x > 0) && (y < 10)
	runTest("2) Logic AND: (x > 0) && (y < 10)", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		y := symbolic.NewSymbolicVariable("y", symbolic.IntType)
		zero := symbolic.NewIntConstant(0)
		ten := symbolic.NewIntConstant(10)

		cond := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{
			symbolic.NewBinaryOperation(x, zero, symbolic.GT),
			symbolic.NewBinaryOperation(y, ten, symbolic.LT),
		}, symbolic.AND)
		tr.AssertExpr(cond)

		if m := mustSat(tr); m != nil {
			xv, _ := tr.GetIntValue(m, x)
			yv, _ := tr.GetIntValue(m, y)
			fmt.Printf("Model: x=%d, y=%d\n", xv, yv)
		}
	})

	// 3) Арифметика без unary: (x - 5) == 0  => x == 5
	runTest("3) Arithmetic (no unary): (x - 5) == 0", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		sub := symbolic.NewBinaryOperation(x, symbolic.NewIntConstant(5), symbolic.SUB)
		eq0 := symbolic.NewBinaryOperation(sub, symbolic.NewIntConstant(0), symbolic.EQ)
		tr.AssertExpr(eq0)

		if m := mustSat(tr); m != nil {
			xv, _ := tr.GetIntValue(m, x)
			fmt.Printf("Model: x=%d (expect 5)\n", xv)
		}
	})

	// 4) NOT без unary:  !(x > 0)
	runTest("4) NOT: !(x > 0)", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		zero := symbolic.NewIntConstant(0)
		gt := symbolic.NewBinaryOperation(x, zero, symbolic.GT)
		not := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{gt}, symbolic.NOT)
		tr.AssertExpr(not)

		if m := mustSat(tr); m != nil {
			xv, _ := tr.GetIntValue(m, x)
			fmt.Printf("Model: x=%d (should be <= 0)\n", xv)
		}
	})

	// 5) Arrays (Int): A' = store(A, 5, 42);  A'[5] == 42
	runTest("5) Arrays Int: store/select A' = store(A, 5, 42)", func(tr *translator.Z3Translator) {
		A := symbolic.NewArray("A", symbolic.IntType)
		Aprime := symbolic.NewArray("A_prime", symbolic.IntType)
		i := symbolic.NewIntConstant(5)
		val := symbolic.NewIntConstant(42)

		tr.AssertExpr(symbolic.NewBinaryOperation(
			Aprime,
			symbolic.NewArrayStore(A, i, val),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			symbolic.NewArrayAccess(Aprime, i),
			val,
			symbolic.EQ,
		))

		if m := mustSat(tr); m != nil {
			term, _ := tr.TranslateExpression(symbolic.NewArrayAccess(Aprime, i))
			v := m.Eval(term.(z3.Value), true)
			fmt.Printf("Model: A'[5]=%s (expect 42)\n", v.String())
		}
	})

	// 6) Arrays (Bool): B' = store(B, 7, true);  B'[7] == true
	runTest("6) Arrays Bool: store/select", func(tr *translator.Z3Translator) {
		B := symbolic.NewArray("B", symbolic.BoolType)
		Bp := symbolic.NewArray("B_prime", symbolic.BoolType)
		j := symbolic.NewIntConstant(7)
		trueC := symbolic.NewBoolConstant(true)

		tr.AssertExpr(symbolic.NewBinaryOperation(
			Bp,
			symbolic.NewArrayStore(B, j, trueC),
			symbolic.EQ,
		))
		tr.AssertExpr(symbolic.NewBinaryOperation(
			symbolic.NewArrayAccess(Bp, j),
			trueC,
			symbolic.EQ,
		))
		_ = mustSat(tr)
	})

	// 7) Arrays equality/inequality: A != C (оба Int-массивы)
	runTest("7) Arrays equality/inequality: A != C", func(tr *translator.Z3Translator) {
		A := symbolic.NewArray("A", symbolic.IntType)
		C := symbolic.NewArray("C", symbolic.IntType)
		tr.AssertExpr(symbolic.NewBinaryOperation(A, C, symbolic.NE))
		_ = mustSat(tr)
	})

	// 8) FunctionCall: Bool и IMPLIES
	runTest("8) FunctionCall Bool + IMPLIES", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		y := symbolic.NewSymbolicVariable("y", symbolic.IntType)

		Pxy := symbolic.NewFunctionCall("P", []symbolic.SymbolicExpression{x, y}, symbolic.BoolType)
		P00 := symbolic.NewFunctionCall("P", []symbolic.SymbolicExpression{symbolic.NewIntConstant(0), symbolic.NewIntConstant(0)}, symbolic.BoolType)

		// Ассерты: P(x,y)  и  (P(0,0) => P(x,y))
		tr.AssertExpr(Pxy)
		tr.AssertExpr(symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{P00, Pxy}, symbolic.IMPLIES))

		if m := mustSat(tr); m != nil {
			fmt.Printf("Model = %s\n", m.String())
		}
	})

	// 9) FunctionCall Int + ITE: max(x,y) = ite(x>y,x,y), max >= x && max >= y
	// и отдельный UNSAT-кейс: x>0 && x<0
	runTest("9) Conditional (ITE) with Int FunctionCall-style args", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		y := symbolic.NewSymbolicVariable("y", symbolic.IntType)

		cond := symbolic.NewBinaryOperation(x, y, symbolic.GT)
		maxXY := symbolic.NewConditionalExpression(cond, x, y)

		tr.AssertExpr(symbolic.NewBinaryOperation(maxXY, x, symbolic.GE))
		tr.AssertExpr(symbolic.NewBinaryOperation(maxXY, y, symbolic.GE))
		if m := mustSat(tr); m != nil {
			fmt.Printf("Model = %s\n", m.String())
		}
	})

	runTest("10) Contradiction: x > 0 && x < 0 (UNSAT)", func(tr *translator.Z3Translator) {
		x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
		zero := symbolic.NewIntConstant(0)
		tr.AssertExpr(symbolic.NewBinaryOperation(x, zero, symbolic.GT))
		tr.AssertExpr(symbolic.NewBinaryOperation(x, zero, symbolic.LT))
		_ = mustSat(tr) // выведет UNSAT
	})
}
