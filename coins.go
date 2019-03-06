package main

import (
	"math"
	"fmt"
)

func getWays(n int64, c []int64) int64 {
	return coins_dp(c, len(c)-1, n)
}

func coins_dp(coins []int64, n int, cap int64) int64 {
	var result int64
	if n < 0 || coins[n] > cap {
		return 0
	} else {
		t1 := coins[n] + coins_dp(coins, n-1, cap-coins[n])
		t2 := coins_dp(coins, n-1, cap)
		result = int64(math.Max(float64(t1), float64(t2)))
	}

	return result
}
func main() {
	fmt.Println(getWays(10, []int64{1,9}))
}
