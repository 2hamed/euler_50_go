package main

import "fmt"

func shop(n int32, k int32, centers []string, roads [][]int32) int32 {
	adj := make([][]int32, n)

	for i := int32(0); i < n; i++ {
		adj[i] = make([]int32, n)
	}

	for _, road := range roads {
		adj[road[0]-1][road[1]-1] = road[2]
	}

	return 0
}

func main() {
	centers := []string{
		"1 1",
		"1 2",
		"1 3",
		"1 4",
		"1 5",
	}

	roads := [][]int32{
		[]int32{1, 2, 10},
		[]int32{1, 3, 10},
		[]int32{2, 4, 10},
		[]int32{3, 5, 10},
		[]int32{4, 5, 10},
	}
	minTime := shop(5, 5, centers, roads)

	fmt.Println("minTime:", minTime)
}
