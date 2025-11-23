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

func OnStackInt() *int {
	var foo int

	foo = 5

	return &foo
}

func OnHeapInt() *int {
	foo := new(int)

	*foo = 5 

	return foo
}

func ParamArg(alias1 *int, alias2 *int) *int {

	return alias1
}
`

	// Создаём builder для SSA
	builder := ssa.NewBuilder()

	// Строим SSA из исходного кода
	fn, err := builder.ParseAndBuildSSA(source, "ParamArg")
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(fn.Blocks))

	fmt.Println("\n=== SSA dump функции testStructBasic ===")
	var buf bytes.Buffer
	gosssa.WriteFunction(&buf, fn) // WriteFunction ждёт *bytes.Buffer
	buf.WriteTo(os.Stdout)         // выводим на stdout

	dumpStructTypes(fn)
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
