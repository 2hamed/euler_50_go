package main

import "fmt"

func main() {
	a := []int{3, 6, 1, 4, 9, 5, 6, 11, 0}

	for i := 1; i < len(a); i++ {
		t := a[i]
		j := i
		for j > 0 && a[j-1] > t {
			a[j], a[j-1] = a[j-1], a[j]
			j--
		}
	}

	fmt.Println(a)
}
