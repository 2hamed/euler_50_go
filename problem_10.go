package main

import (
	"fmt"
	"math"
)

func main() {
	var sum int64 = 0
	for i := int64(2); i <= 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	fmt.Println("Sum: ", sum)
}
func isPrime(i int64) bool {
	s := math.Sqrt(float64(i))
	for j := int64(2); float64(j) <= s; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}