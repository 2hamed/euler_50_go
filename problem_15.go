package main

import "fmt"

func main() {
	pathcount := make(map[int]int)
	fmt.Println(count(0, 21, pathcount))
}

func count(i int, n int, pathCount map[int]int) int {
	a, b := adjacent(i, n)
	//fmt.Printf("adjacent of: %d --> %d, %d\n", i, a, b)
	if c, ok := pathCount[i]; ok {
		return c
	}

	if a == -1 && b == -1 {
		return 1
	}
	if i == n*n-n || i == n*n-1 {
		return 1
	}
	ca := 0
	if a != -1 {
		ca = count(a, n, pathCount)
		pathCount[a] = ca

	}
	cb := 0
	if b != -1 {
		cb = count(b, n, pathCount)
		pathCount[b] = cb
	}
	return ca + cb
}

func adjacent(i int, n int) (int, int) {
	if i >= (n*n)-1 {
		return -1, -1
	}
	if i%n == n-1 {
		return i + n, -1
	}
	if i/n == n-1 {
		return -1, i + 1
	}
	return i + 1, i + n
}