// Демонстрационная программа для тестирования SSA построения
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"symbolic-execution-course/internal/ssa"

	"go/types"

	gosssa "golang.org/x/tools/go/ssa"
)

func dumpStructTypes(fn *gosssa.Function) {
	if fn.Pkg == nil {
		fmt.Println("Функция не принадлежит пакету (nil Pkg)")
		return
	}
	tpkg := fn.Pkg.Pkg // *types.Package
	scope := tpkg.Scope()

	fmt.Println("\n=== Типы-структуры в пакете ===")

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		tn, ok := obj.(*types.TypeName)
		if !ok {
			continue
		}

		// Интересуют только те, у кого underlying — struct
		if st, ok := tn.Type().Underlying().(*types.Struct); ok {
			fmt.Printf("type %s struct {\n", tn.Name())
			for i := 0; i < st.NumFields(); i++ {
				f := st.Field(i)
				fmt.Printf("    %s %s\n", f.Name(), f.Type().String())
			}
			fmt.Println("}")
		}
	}
}

func main() {
	fmt.Println("=== SSA Builder Demo ===")

	// Пример исходного кода для анализа
	source := `
package main

type Person struct {
	Name string
	Age  int
	ID   int
}

func testStructPointer() *Person {
	p := &Person{Name: "Bob", Age: 30, ID: 2002}
	p.Age = p.Age + 5
	return p
}

func testStructPointerModification(p *Person) {
	if p != nil {
		p.Age = p.Age + 10
		p.ID = p.ID + 1000
	}
}

func testNestedStructPointer() {
	p := testStructPointer()
	testStructPointerModification(p)
}

func testArrayFixed() [5]int {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i * i
	}
	return arr
}

func testArrayModification(arr [5]int) [5]int {
	for i := range arr {
		arr[i] = arr[i] + 1
	}
	return arr
}

func testSliceCreation() []int {
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = i * 2
	}
	return slice
}

func testSliceAppend() []int {
	var slice []int
	for i := 0; i < 3; i++ {
		slice = append(slice, i*10)
	}
	return slice
}

func testSliceModification(slice []int) []int {
	for i := range slice {
		slice[i] = slice[i] * 2
	}
	slice = append(slice, 999)
	return slice
}

type Student struct {
	Name    string
	Grades  [5]int
	Average float64
}

func testStructWithArray() Student {
	var s Student
	s.Name = "Charlie"
	s.Grades = [5]int{85, 90, 78, 92, 88}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Average = float64(sum) / float64(len(s.Grades))

	return s
}
`

	// Создаём builder для SSA
	builder := ssa.NewBuilder()

	// Строим SSA из исходного кода
	fn, err := builder.ParseAndBuildSSA(source, "testArrayModification")
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(fn.Blocks))

	fmt.Println("\n=== SSA dump функции testArrayModification ===")
	var buf bytes.Buffer
	gosssa.WriteFunction(&buf, fn) // WriteFunction ждёт *bytes.Buffer
	buf.WriteTo(os.Stdout)         // выводим на stdout

	dumpStructTypes(fn)
}
