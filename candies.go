package main

import "fmt"

func candies(n int32, arr []int32) int64 {
	s := candy(n, arr)
	arev := make([]int32, len(arr))
	copy(arev, arr)
	for i := len(arev)/2 - 1; i >= 0; i-- {
		opp := len(arev) - 1 - i
		arev[i], arev[opp] = arev[opp], arev[i]
	}
	s2 := candy(n, arev)
	if s2 < s {
		return s2
	}
	return s
}
func candy(n int32, arr []int32) int64 {
	candiess := make([]int, 0)
	for i := range arr {
		if i == 0 {
			if arr[0] > arr[1] {
				candiess = append(candiess, 2)
			} else {
				candiess = append(candiess, 1)
			}
		} else {
			if arr[i] > arr[i-1] {
				candiess = append(candiess, candiess[i-1]+1)
			} else {
				if i+1 >= int(n) {
					candiess = append(candiess, 1)
				} else {
					t := candiess[i-1]
					t2 := 0
					for j := i; j < int(n); j++ {
						if arr[j] < arr[j-1] {
							t2++
						} else {
							break
						}
					}
					if t2 >= t {
						candiess[i-1] = t2 + 1
					}
					candiess = append(candiess, t2)
				}
			}
		}
	}
	//fmt.Println(candiess)
	sum := 0
	for _, v := range candiess {
		sum += v
	}
	return int64(sum)
}

func main() {
	fmt.Println(candies(8, []int32{2, 4, 3, 5, 2, 6, 4, 5}))
}
