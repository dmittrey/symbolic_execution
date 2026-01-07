package internal

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal/ssa"
	"symbolic-execution-course/internal/translator"
)

type Analyser struct {
	SSABuilder   *ssa.Builder
	StatesQueue  PriorityQueue
	PathSelector PathSelector
	Results      []Interpreter
	Z3Translator *translator.Z3Translator
}

func NewAnalyser() Analyser {
	return Analyser{
		SSABuilder:   ssa.NewBuilder(),
		Z3Translator: translator.NewZ3Translator(),
	}
}

func (an *Analyser) Analyse(source string, functionName string) []Interpreter {
	// Создаём builder для SSA
	builder := ssa.NewBuilder()

	// Строим SSA из исходного кода
	fn, err := builder.ParseAndBuildSSA(source, functionName)
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(fn.Blocks))
}
