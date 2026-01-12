package internal

import "math/rand"

type PathSelector interface {
	CalculatePriority(interpreter Interpreter) int
}

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

type DfsPathSelector struct {
	counter int
}

func NewDfsPathSelector() *DfsPathSelector {
	return &DfsPathSelector{counter: MinInt}
}

func (dfs *DfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	dfs.counter++
	return dfs.counter
}

type BfsPathSelector struct {
	counter int
}

func NewBfsPathSelector() *BfsPathSelector {
	return &BfsPathSelector{counter: MaxInt}
}

func (bfs *BfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	bfs.counter--
	return bfs.counter
}

type RandomPathSelector struct{}

func NewRandomPathSelector() *RandomPathSelector { return &RandomPathSelector{} }

func (random *RandomPathSelector) CalculatePriority(interpreter Interpreter) int {
	return rand.Int()
}
