package main

import "fmt"

func rotLeft(a []int32, d int32) []int32 {
	t := a
	c := len(a)
	b := make([]int32, c)
	dd := int(d) % c

	for j := 0; j < c-dd; j++ {
		b[j] = a[j+dd]
	}
	j := 0
	for i := c - dd; i < c; i++ {
		b[i] = t[i-(c-dd)]
		j++
	}

	return b
}

func main() {
	fmt.Println(rotLeft([]int32{33, 47, 70, 37, 8, 53, 13, 93, 71, 72, 51, 100, 60, 87, 97}, 15252))
}
