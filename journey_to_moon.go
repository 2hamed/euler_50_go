package main

import "fmt"

func journeyToMoon(n int32, astronaut [][]int32) int32 {
	adj := make(map[int32][]int32)

	for _, a := range astronaut {
		if _, ok := adj[a[0]]; !ok {
			adj[a[0]] = make([]int32, 0)
		}
		adj[a[0]] = append(adj[a[0]], a[1])

		if _, ok := adj[a[1]]; !ok {
			adj[a[1]] = make([]int32, 0)
		}
		adj[a[1]] = append(adj[a[1]], a[0])
	}

	visited := make(map[int32]bool)

	for i := n - 1; i != 0; i-- {
		if _, ok := adj[i]; !ok {
			adj[i] = make([]int32, 0)
		}
		visited[i] = false
	}

	components := make(map[int32]int32)

outer:
	for {
		for k, v := range visited {
			if !v {
				path := make([]int32, 0)
				path = dfs(k, visited, adj, path)
				components[k] = int32(len(path))
				continue outer
			}
		}
		break
	}

	// fmt.Println(components)

	var sum int32
	var result int32
	for _, size := range components {
		result += sum * size
		sum += size
	}

	return result

}

func dfs(current int32, visited map[int32]bool, adj map[int32][]int32, path []int32) []int32 {
	if v := visited[current]; v {
		return path
	}

	visited[current] = true
	path = append(path, current)
	for _, a := range adj[current] {
		path = dfs(a, visited, adj, path)
	}
	return path
}

func main() {
	a := [][]int32{
		[]int32{1, 2},
		[]int32{3, 4},
	}
	s := journeyToMoon(10000, a)

	fmt.Println(s)
}
