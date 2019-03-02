package main

import (
	"fmt"
)

func main() {
	var max float32
	calculated := make(map[int]bool)
	for i := 1000000; i > 1; i-- {
		nn := float32(i) / float32(phi(i))
		if nn > max {
			max = nn
		}
	}

	fmt.Println(max)
}

func phi(n int) int {
	var count int
	for i := n - 1; i > 0; i++ {
		if n%i != 0 {
			count++
		}
	}
	return count
}
