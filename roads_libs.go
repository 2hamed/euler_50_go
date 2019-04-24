package main

import (
	"fmt"
)

// Complete the roadsAndLibraries function below.
func roadsAndLibraries(n int32, c_lib int32, c_road int32, cities [][]int32) int64 {
	if c_lib <= c_road {
		return int64(c_lib) * int64(n)
	}
	adj := make(map[int32][]int32)
	for _, c := range cities {
		if _, ok := adj[c[0]]; !ok {
			adj[c[0]] = make([]int32, 0)
		}
		adj[c[0]] = append(adj[c[0]], c[1])

		if _, ok := adj[c[1]]; !ok {
			adj[c[1]] = make([]int32, 0)
		}
		adj[c[1]] = append(adj[c[1]], c[0])
	}

	visited := make(map[int32]bool)

	for i := int32(1); i <= n; i++ {
		visited[i] = false
	}
	comps := make([]int32, 0)
outer:
	for {
		for k, v := range visited {
			if !v {
				path := make([]int32, 0)
				path = dfs(k, path, visited, adj)
				comps = append(comps, int32(len(path)))
				continue outer
			}
		}
		break
	}
	cost := int64(0)
	for i := 0; i < len(comps); i++ {
		cost += int64(comps[i]-1)*int64(c_road) + int64(c_lib)
	}
	return cost
}

func dfs(cur int32, path []int32, visited map[int32]bool, adj map[int32][]int32) []int32 {
	if v := visited[cur]; v {
		return path
	}

	visited[cur] = true

	path = append(path, cur)

	for _, v := range adj[cur] {
		path = dfs(v, path, visited, adj)
	}
	return path
}

func main() {
	cities := make([][]int32, 4)
	cities[0] = []int32{1, 2}
	cities[1] = []int32{2, 3}
	cities[2] = []int32{1, 3}
	cities[3] = []int32{4, 5}

	c := roadsAndLibraries(5, 1, 1, cities)
	fmt.Println(c)
}
