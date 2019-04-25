package main

import "fmt"

func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	adj := make([][]int32, n)
	for i := range adj {
		adj[i] = make([]int32, n)
	}

	for _, e := range edges {
		adj[e[0]-1][e[1]-1] = 1
		adj[e[1]-1][e[0]-1] = 1
	}

	for _, c := range adj {
		fmt.Println(c)
	}

	fin := make([]int32, 0)

	for i := int32(0); i < n; i++ {
		if i == s-1 {
			continue
		}
		path := make([]int32, 0)
		path = bfSearch(s-1, i, adj, path)
		fmt.Println(s-1, i, path)
		if path[0] == -1 {
			fin = append(fin, -1)
		} else if len(path) > 2 {
			fin = append(fin, (int32(len(path)) - 1))
		} else {
			fin = append(fin, int32(len(path)))
		}
	}

	return fin
}

func bfSearch(start, end int32, adj [][]int32, path []int32) []int32 {
	if start == end {
		return []int32{-1}
	}
	for e, p := range adj[start] {
		if int32(e) == end && p == 1 {
			return []int32{1}
		}
	}

	for i, v := range adj[start] {
		if v == 1 {
			path = append(path, bfSearch(int32(i), end, adj, path)...)
		}
	}

	return append(path, 1)
}

func main() {
	edges := [][]int32{
		[]int32{1, 2},
		[]int32{1, 3},
		[]int32{2, 4},
		[]int32{4, 5},
		[]int32{5, 6},
	}

	fmt.Println(bfs(20, 2, edges, 10))
}
