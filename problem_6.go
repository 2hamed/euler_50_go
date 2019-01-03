package main

import (
	"fmt"
	"math"
)

func main() {
	var sqauredSum int64
	var sumSquared int64

	for i := int64(1); i <= 100; i++ {
		sqauredSum += i * i
		sumSquared += i
	}
	sumSquared *= sumSquared

	fmt.Println("Diff: ", int64(math.Abs(float64(sumSquared-sqauredSum))))
}
