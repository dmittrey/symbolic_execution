package internal

import (
	"container/heap"
	"context"
	"fmt"
	"log"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	const maxVisitsPerLoc = 5

	builder := ssa.NewBuilder()

	fn, err := builder.ParseAndBuildSSA(source, functionName)
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(fn.Blocks))

	selector := NewRandomPathSelector()

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startState := NewInterpreter(analyser, symbolic.NewBoolConstant(true))

	heap.Push(&pq, &Item{
		value:    startState,
		block:    fn.Blocks[0],
		instrIdx: 0,
		priority: selector.CalculatePriority(startState),
		pred:     nil,
		visits:   make(map[LocKey]uint8),
	})

	results := make([]Interpreter, 0)

	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*Item)

		if it.block == nil {
			results = append(results, it.value)
			continue
		}

		select {
		case <-ctx.Done():
			fmt.Printf("Analyse stopped: %v\n", ctx.Err())
			return results
		default:
		}

		loc := LocKey{Block: it.block.Index, InstrIx: it.instrIdx}
		cnt := it.visits[loc]
		if cnt >= maxVisitsPerLoc {
			continue
		}

		parentVisits := cloneVisits(it.visits)
		parentVisits[loc] = cnt + 1

		st := it.value
		nextItems := (&st).interpretDynamically(it.block, it.instrIdx, it.pred)

		for _, ni := range nextItems {
			ni.visits = cloneVisits(parentVisits) // <-- прокидываем
			ni.priority = selector.CalculatePriority(ni.value)
			heap.Push(&pq, ni)
		}
	}

	return results
}

func cloneVisits(src map[LocKey]uint8) map[LocKey]uint8 {
	if src == nil {
		return make(map[LocKey]uint8)
	}
	dst := make(map[LocKey]uint8, len(src)+1)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
