package operator

import "fmt"

type MemcachedOperator struct {
	nodes []Node
	min   int
	max   int
}

type Node struct {
	id      string
	healthy bool
}

func NewMemcachedOperator(min, max int) *MemcachedOperator {
	nodes := make([]Node, min)
	for i := 0; i < min; i++ {
		nodes[i] = Node{id: fmt.Sprintf("node%d", i), healthy: true}
	}

	return &MemcachedOperator{
		min:   min,
		max:   max,
		nodes: nodes,
	}
}
