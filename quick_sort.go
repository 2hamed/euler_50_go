package main

import (
	"fmt"
	"math/rand"
)

func quickSort(a []int, left, right int) {
	if left < right {
		pivotLocation := partition(a, left, right)
		//fmt.Println("Pivot:", pivotLocation)
		quickSort(a, left, pivotLocation-1)
		quickSort(a, pivotLocation+1, right)
	}
}

func partition(a []int, left, right int) int {
	p := a[right]

	var i, j = left, right-1
	for {
		for i <= j && a[i] <= p {
			i++
		}

		for j >= i && a[j] >= p {
			j--
		}
		if j < i {
			break
		} else {
			a[i], a[j] = a[j], a[i]
		}
	}
	a[right], a[i] = a[i], a[right]

	return i
}

func main() {
	a := func() []int {
		a := make([]int, 0)
		for len(a) < 1000000 {
			a = append(a, rand.Intn(1000000))
		}
		return a
	}()

	quickSort(a, 0, len(a)-1)

	fmt.Println(len(a))
}
