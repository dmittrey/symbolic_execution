package internal

import (
	"math"
	"math/rand"
)

type PathSelector interface {
	CalculatePriority(interpreter Interpreter) int
}

type DfsPathSelector struct {
	counter int // int.min_value
}

func NewDfsPathSelector() *DfsPathSelector {
	return &DfsPathSelector{counter: math.MinInt64}
}

func (dfs *DfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	dfs.counter++
	return dfs.counter
}

type BfsPathSelector struct {
	counter int // int.max_value
}

func NewBfsPathSelector() *BfsPathSelector {
	return &BfsPathSelector{counter: math.MaxInt64}
}

func (bfs *BfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	bfs.counter--
	return bfs.counter
}

type RandomPathSelector struct{}

func (random *RandomPathSelector) CalculatePriority(interpreter Interpreter) int {
	return rand.Int()
}
