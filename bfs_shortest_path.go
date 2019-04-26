package main

import "fmt"

func bfs(n int32, m int32, edges [][]int32, start int32) []int32 {
	adj := make([][]int32, n)
	for i := range adj {
		adj[i] = make([]int32, n)
	}

	for _, e := range edges {
		adj[e[0]-1][e[1]-1] = 1
		adj[e[1]-1][e[0]-1] = 1
	}

	for i, c := range adj {
		fmt.Println(i+1, c)
	}

	result := make([]int32, 0)

	for j := int32(0); j < n; j++ {
		if j == start-1 {
			continue
		}
		q := make([]int32, 0)

		levels := make([]int, n)
		visited := make([]bool, n)

		t := int32(j)

		q = append(q, start-1)

		found := false
		for len(q) > 0 {
			s := q[0]
			q = q[1:]
			visited[s] = true
			if s == t {
				found = true
				break
			}
			for i, v := range adj[s] {
				if v == 1 {
					if vis := visited[i]; !vis {
						q = append(q, int32(i))
					}
				}
			}
			// fmt.Println(q)
		}
		f := "found"
		if !found {
			f = "not found"
		}
		fmt.Printf("from %d to %d %s in %d steps\n", start, j+1, f, count)

		// fmt.Println(found, count)
		if found {
			result = append(result, count*6)
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
