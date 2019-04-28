package main

import (
	"container/heap"
	"fmt"
)

func main() {
	edges := [][]int32{
		[]int32{1, 2, 24},
		[]int32{1, 4, 20},
		[]int32{3, 1, 3},
		[]int32{4, 3, 12},
		[]int32{4, 1, 2},
	}
	fmt.Println(shortestReach(4, edges, 1))
}

const MaxInt int32 = int32(^uint32(0) >> 1)

func shortestReach(n int32, edges [][]int32, s int32) []int32 {
	dists := make([]int32, n)

	for i := int32(0); i < n; i++ {
		dists[i] = MaxInt
	}

	adj := make([][]int32, n)

	visited := make([]bool, n)

	for i := int32(0); i < n; i++ {
		adj[i] = make([]int32, n)
	}

	for _, e := range edges {
		if v := adj[e[0]-1][e[1]-1]; v == 0 || e[2] <= v {
			adj[e[0]-1][e[1]-1] = e[2]
		}
		if v := adj[e[1]-1][e[0]-1]; v == 0 || e[2] <= v {
			adj[e[1]-1][e[0]-1] = e[2]
		}
	}

	dists[s-1] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &NW{Node: s - 1, Weight: 0})

	for pq.Len() != 0 {
		p := heap.Pop(&pq).(*NW)
		if visited[p.Node] {
			continue
		}
		visited[p.Node] = true
		for i, w := range adj[p.Node] {
			if w > 0 && dists[p.Node]+w < dists[i] {
				dists[i] = dists[p.Node] + w
				heap.Push(&pq, &NW{Node: int32(i), Weight: dists[i]})
			}
		}
	}

	finDists := make([]int32, 0)

	for i, d := range dists {
		if int32(i) != s-1 {
			if d == MaxInt {
				d = -1
			}
			finDists = append(finDists, d)
		}
	}

	return finDists
}

type NW struct {
	Node   int32
	Weight int32
}

type PriorityQueue []*NW

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*NW)

	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *NW, node int32, weight int32) {
	item.Node = node
	item.Weight = weight
	// heap.Fix(pq, item.Index)
}
