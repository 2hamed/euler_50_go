package main

import (
	"container/heap"
	"fmt"
)

func main() {
	pq := PriorityQueue{
		&NW{Node: 1, Weight: 18, Index: 0},
		&NW{Node: 2, Weight: 2, Index: 1},
		&NW{Node: 3, Weight: 18, Index: 2},
		&NW{Node: 4, Weight: 1, Index: 3},
	}

	heap.Init(&pq)
	nw := heap.Pop(&pq).(*NW)
	fmt.Printf("%d:%d --> %d\n", nw.Node, nw.Weight, nw.Index)

	nw = heap.Pop(&pq).(*NW)
	fmt.Printf("%d:%d --> %d\n", nw.Node, nw.Weight, nw.Index)
}


type NW struct {
	Node   int32
	Weight int32
	Index  int
}

type PriorityQueue []*NW

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*NW)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *NW, node int32, weight int32) {
	item.Node = node
	item.Weight = weight
	heap.Fix(pq, item.Index)
}