package main

import (
	"fmt"
)

var memo = make(map[int64]map[int]int64)

func getWays(cap int64, c []int64) int64 {
	return getWaysDP(cap, c, len(c))

}

func getWaysDP(capRemaining int64, coins []int64, coinNum int) int64 {
	if capRemaining == 0 {
		return 1
	} else if capRemaining < 0 || coinNum <= 0 {
		return 0
	} else if r, ok := memo[capRemaining][coinNum]; ok {
		return r
	}

	r := getWaysDP(capRemaining, coins, coinNum-1) + getWaysDP(capRemaining-coins[coinNum-1], coins, coinNum)
	if memo[capRemaining] == nil {
		memo[capRemaining] = make(map[int]int64)
	}
	memo[capRemaining][coinNum] = r
	return r
}

func main() {
	fmt.Println(getWays(10, []int64{2, 5, 3, 6}))
}
