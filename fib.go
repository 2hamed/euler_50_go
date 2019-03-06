package main

import "fmt"

func main() {
	dp := make(map[float64]float64)
	fmt.Println(fib(1000, dp))
}

func fib(n float64, dp map[float64]float64) float64 {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	if v, ok := dp[n]; ok {
		return v
	}

	r := fib(n-2, dp) + fib(n-1, dp)
	dp[n] = r
	return r
}
