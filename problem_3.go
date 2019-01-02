package main

import (
	"fmt"
	"math"
)

func main() {
	n := 600851475143
	maxPrime := -1

	for n%2 == 0 {
		maxPrime = 2
		n >>= 1
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			maxPrime = i
			n = n / i
		}
	}
	if n > 2 {
		maxPrime = n
	}

	fmt.Println("Max prime factor: ", maxPrime)
}
