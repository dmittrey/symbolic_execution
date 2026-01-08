package main

import (
	"fmt"
	"symbolic-execution-course/internal"
)

func main() {
	source := `
package main

func testFunction(x int) int {
	if x > 0 {
		return x * 2
	} else {
		return x * -1
	}
}
`
	analyser := internal.NewAnalyser()

	result := analyser.Analyse(source, "testFunction")
	for _, st := range result {
		fmt.Println(st.Dump())
	}
}
