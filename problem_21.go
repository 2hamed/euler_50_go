package main

import (
	"math"
	"fmt"
)

var divSumMap = make(map[int]int)

func main() {
	sum := 0
	for i := 10000; i > 2; i-- {
		if isAmicable(i) {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)
}

func divSum(n int) int {
	if v, ok := divSumMap[n]; ok {
		return v
	}
	sqrt := int(math.Sqrt(float64(n)))
	sum := 0
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			if i == n/i {
				sum += i
			} else {
				sum += i + n/i
			}
		}
	}
	divSumMap[n] = sum + 1
	return sum + 1
}

func isAmicable(n int) bool {
	divS := divSum(n)
	if divS == n {
		return false
	}
	if divSum(divS) == n {
		return true
	}

	return false

}
