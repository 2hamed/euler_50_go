package main

import "fmt"

func bfs(n int32, m int32, edges [][]int32, start int32) []int32 {
	adj := make([][]int32, n)

	for _, e := range edges {
		if l := adj[e[0]-1]; l == nil {
			adj[e[0]-1] = make([]int32, 0)
		}
		adj[e[0]-1] = append(adj[e[0]-1], e[1]-1)

		if l := adj[e[1]-1]; l == nil {
			adj[e[1]-1] = make([]int32, 0)
		}
		adj[e[1]-1] = append(adj[e[1]-1], e[0]-1)
	}

	// for i, c := range adj {
	// fmt.Println(i+1, c)
	// }

	result := make([]int32, 0)

	for j := int32(0); j < n; j++ {
		if j == start-1 {
			continue
		}
		q := make([]int32, 0)

		levels := make([]int32, n)
		visited := make([]bool, n)

		t := int32(j)

		q = append(q, start-1)

		levels[start-1] = 0

		found := false
		for len(q) > 0 {
			s := q[0]
			q = q[1:]
			visited[s] = true
			if s == t {
				found = true
				break
			}
			if adj[s] == nil {
				continue
			}
			for _, v := range adj[s] {
				if vis := visited[v]; !vis {
					levels[v] = levels[s] + 1
					q = append(q, v)
					visited[v] = true
				}
			}
			// fmt.Println(q)
		}

		if found {
			result = append(result, levels[j]*6)
		} else {
			result = append(result, -1)
		}
	}

	return result
}

func main() {
	edges := [][]int32{
		[]int32{1, 2},
		[]int32{1, 3},
		[]int32{3, 4},
	}

	fmt.Println(bfs(5, 3, edges, 1))
}
