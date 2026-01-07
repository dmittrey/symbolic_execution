package internal

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/ssa"
	"symbolic-execution-course/internal/symbolic"
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

func (analyser *Analyser) Analyse(source string, functionName string) []Interpreter {
	// Создаём builder для SSA
	builder := ssa.NewBuilder()

	// Строим SSA из исходного кода
	fn, err := builder.ParseAndBuildSSA(source, functionName)
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(fn.Blocks))

	entry := fn.Blocks[0]
	startState := Interpreter{
		CallStack:     make([]CallStackFrame, 0),
		Analyser:      analyser,
		PathCondition: symbolic.NewBoolConstant(true),
		Heap:          memory.NewSymbolicMemory(),
	}
	traverse_dfs(entry, 0, startState)

	return nil
}
