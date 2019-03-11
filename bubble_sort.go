package main

import "fmt"

func main() {
	a := []int{5, 9, 74, 23, 4, 6, 541, 887, 65, 1, 58, 9, 6}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Println(a)
}
