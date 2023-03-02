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

func (o *MemcachedOperator) Scale(n int) {
	if n < 0 {
		o.scaleDown(-n)
	}

	if n > 0 {
		o.scaleUp(n)
	}
}

func (o *MemcachedOperator) CountNodes() int {
	return len(o.nodes)
}

func (o *MemcachedOperator) scaleDown(n int) {
	for i := 0; i < n; i++ {
		if len(o.nodes) > o.max {
			o.nodes = o.nodes[:len(o.nodes)-1]
		}
	}
}

func (o *MemcachedOperator) scaleUp(n int) {
	for i := 0; i < n; i++ {
		if len(o.nodes) < o.max {
			o.nodes = append(o.nodes, Node{id: fmt.Sprintf("node%d", len(o.nodes)), healthy: true})
		}
	}
}
