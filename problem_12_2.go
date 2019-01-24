package main

import (
	"fmt"
	"math"
)

// we use prime factorization again and the number of divisors of a number is equal to the product of exponents of its prime factors plus one

func main() {
	tNumber := 0
	for i := 1; true; i++ {
		tNumber += i
		f := primeFactors(tNumber)
		fmt.Printf("%d: %d -> %d \n", i, tNumber, f)
		if f > 500 {
			fmt.Println("Number: ", tNumber)
			return
		}
	}
}

func primeFactors(n int) int {
	primes := make([]int, 0)

	pm := make(map[int]int)

	for i := 2; i <= int(math.Sqrt(float64(n))); i += 1 {
		for n%i == 0 {
			primes = append(primes, i)
			n = n / i
			if _, ok := pm[i]; ok {
				pm[i]++
			} else {
				pm[i] = 1
			}
		}
	}

	if n > 2 {
		primes = append(primes, n)
		pm[n] = 1
	}
	sum := 1
	for _, v := range pm {
		sum *= v + 1
	}

	return sum
}
