package main

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal"
	"symbolic-execution-course/internal/translator"

	"github.com/ebukreev/go-z3/z3"
)

func main() {
	source := `
package final_tests

type Foo struct {
	a int
}

func IntAliasing(foo1 int, foo2 int) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func BoolAliasing(foo1 bool, foo2 bool) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func StringAliasing(foo1 string, foo2 string) int {
	if foo2 == foo1 {
		return 4
	}
	return 5
}

func IntRefAliasing(foo1 *int, foo2 *int) int {
	*foo2 = 5

	*foo1 = 2

	if *foo2 == 2 {
		return 4
	}
	return 5
}

func BoolRefAliasing(foo1 *bool, foo2 *bool) int {
	*foo2 = true

	*foo1 = false

	if *foo2 == false {
		return 4
	}
	return 5
}

func StringRefAliasing(foo1 *string, foo2 *string) int {
	*foo2 = "hi"

	*foo1 = "kek"

	if *foo2 == "kek" {
		return 4
	}
	return 5
}

func StructAliasing(foo1 *Foo, foo2 *Foo) int {
	foo2.a = 5
	foo1.a = 2
	if foo2.a == 2 {
		return 4
	}
	return 5
}

func ArrayAliasing(arr1 []int, arr2 []int) int {
	arr2[1] = 5
	arr1[1] = 2
	if arr2[1] == 2 {
		return 4
	}
	return 5
}

`

	funcs := []string{
		"IntAliasing",
		"BoolAliasing",
		"StringAliasing",
		"IntRefAliasing",
		"BoolRefAliasing",
		"StringRefAliasing",
		"StructAliasing",
		"ArrayAliasing",
	}

	analyser := internal.NewAnalyser()

	for _, fn := range funcs {
		fmt.Printf("\n==================== %s ====================\n", fn)
		result := analyser.Analyse(source, fn)

		fmt.Printf("States: %d\n", len(result))
		for i, st := range result {
			tr := translator.NewZ3Translator()
			defer tr.Close()

			fmt.Printf("---- State #%d ----\n", i)
			fmt.Println(st.Dump())

			tr.AssertExpr(st.PathCondition)
			for _, c := range st.Heap.Constraints() {
				tr.AssertExpr(c)
			}

			if m := mustSat(tr); m != nil {
				fmt.Printf("Model:\n%s\n", m.String())
			}
		}
	}
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
