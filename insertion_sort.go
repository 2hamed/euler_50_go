package main

import "fmt"

func main() {
	a := []int{9, 8, 4, 5, 1, 7, 3}

	for i := 1; i < len(a); i++ {
		j := i
		for j > 0 && a[j-1] > a[j] {
			a[j], a[j-1] = a[j-1], a[j]
			j--
			fmt.Println(a)
		}
	}

	fmt.Println(a)
}
