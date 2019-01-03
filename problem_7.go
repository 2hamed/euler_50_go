package main

import (
	"math"
	"fmt"
)

func main() {
	var j int
	for i := 2; true; i++ {
		if isPrime(i) {
			j++
		}
		if j == 10001 {
			fmt.Println(i)
			return
		}
	}
}

func isPrime(i int) bool {
	s := math.Sqrt(float64(i))
	for j := 2; float64(j) <= s; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
