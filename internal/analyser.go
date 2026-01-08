package internal

import (
	"container/heap"
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

	// создать пустую очередь
	selector := NewDfsPathSelector()

	// инициализировать как heap
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Подготовить Item
	startState := Interpreter{
		CallStack:     make([]CallStackFrame, 0),
		Analyser:      analyser,
		PathCondition: symbolic.NewBoolConstant(true),
		Heap:          memory.NewSymbolicMemory(),
	}

	// Запушить Item
	heap.Push(&pq, &Item{
		value:    startState,
		block:    fn.Blocks[0],
		instrIdx: 0,
		priority: selector.CalculatePriority(startState),
	})

	results := make([]Interpreter, 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)

		// finished?
		if it.block == nil {
			results = append(results, it.value)
			continue
		}

		st := it.value
		nextItems := (&st).interpretDynamically(it.block, it.instrIdx)

		for _, ni := range nextItems {
			ni.priority = selector.CalculatePriority(ni.value)
			heap.Push(&pq, ni)
		}
	}

	return results
}
