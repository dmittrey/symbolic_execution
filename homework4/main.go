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

// ==== helper types for tests below ====

type Foo struct {
	a int
}

type ObjectWithPrimitivesClass struct {
	x int
	y int
}

func NewObjectWithPrimitivesClass() *ObjectWithPrimitivesClass {
	return &ObjectWithPrimitivesClass{}
}

// ==== aliasing tests ====

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

// ==== arrays tests ====

func DefaultBooleanValues() []bool {
	array := make([]bool, 3)
	if !array[1] {
		return array
	}
	return array
}

func ByteArray(a []byte, x byte) byte {
	if len(a) != 2 {
		return 255
	}
	a[0] = 5
	a[1] = x
	if a[0]+a[1] > 20 {
		return 1
	}
	return 0
}

func CharSizeAndIndex(a []rune, x rune) byte {
	if a == nil || len(a) <= int(x) || x < 1 {
		return 255
	}
	b := make([]rune, x)
	b[0] = 5
	a[x] = x
	if b[0]+a[x] > 7 {
		return 1
	}
	return 0
}

func BooleanArray(arr []bool) int {
	if len(arr) == 0 {
		return 1
	}
	if arr[0] {
		return 2
	}
	arr[0] = true
	return 3
}

func CreateArray(x, y, length int) []*ObjectWithPrimitivesClass {
	if length < 3 {
		return nil
	}

	array := make([]*ObjectWithPrimitivesClass, length)

	for i := 0; i < len(array); i++ {
		array[i] = NewObjectWithPrimitivesClass()
		array[i].x = x + i
		array[i].y = y + i
	}

	return array
}

func IsIdentityMatrix(matrix [][]int) bool {
	if len(matrix) < 3 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) != len(matrix) {
			return false
		}
		for j := 0; j < len(matrix[i]); j++ {
			if i == j && matrix[i][j] != 1 {
				return false
			}
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func ReallyMultiDimensionalArray(array [][][]int) [][][]int {
	if array[1][2][3] != 12345 {
		array[1][2][3] = 12345
	} else {
		array[1][2][3] -= 12345 * 2
	}
	return array
}

func FillMultiArrayWithArray(value []int) [][]int {
	if len(value) < 2 {
		return make([][]int, 0)
	}
	for i := range value {
		value[i] += i
	}
	length := 3
	array := make([][]int, length)
	for i := 0; i < length; i++ {
		array[i] = value
	}
	return array
}
`

	funcs := []string{
		// aliasing
		"IntAliasing",
		"BoolAliasing",
		"StringAliasing",
		"IntRefAliasing",
		"BoolRefAliasing",
		"StringRefAliasing",
		"StructAliasing",
		"ArrayAliasing",

		// arrays
		"DefaultBooleanValues",
		"ByteArray",
		"CharSizeAndIndex",
		"BooleanArray",
		"CreateArray",
		"IsIdentityMatrix",
		"ReallyMultiDimensionalArray",
		"FillMultiArrayWithArray",
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
