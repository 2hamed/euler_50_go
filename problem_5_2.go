package main

import (
	"math"
	"fmt"
)

func main() {
	iMap := make(map[int]int)

	for i := 1; i <= 20; i++ {
		factors := primeFactors(i)

		tMap := make(map[int]int)
		for _, v := range factors {
			if _, ok := tMap[v]; ok {
				tMap[v]++
			} else {
				tMap[v] = 1
			}
		}
		for k, v := range tMap {
			if _, ok := iMap[k]; ok {
				if iMap[k] < v {
					iMap[k] = v
				}
			} else {
				iMap[k] = v
			}
		}
	}

	var product = 1
	for k, v := range iMap {
		product *= int(math.Pow(float64(k), float64(v)))
	}

	fmt.Println("Product: ", product)
}

func primeFactors(n int) []int {
	primes := make([]int, 0)

	for i := 2; i <= int(math.Sqrt(float64(n))); i += 1 {
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
