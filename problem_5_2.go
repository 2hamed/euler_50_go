package main

import (
	"math"
	"fmt"
)

func main() {
	for i := 2; i <= 20; i++ {
		fmt.Println(primeFactors(i))
	}
}

func primeFactors(n int) []int {
	primes := make([]int, 0)
	for n%2 == 0 {
		n >>= 1
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
		}
	}

	if n > 2 {
		primes = append(primes, n)
	}
	return primes
}
