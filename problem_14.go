package main

import "fmt"

func main() {
	max := 0
	var item int64
	for i := int64(999999); i > 0; i-- {
		c := collatz(i)
		if c > max {
			max = c
			item = i
		}
	}

	fmt.Println("Longest: ", item, " Max: ", max)
}

func collatz(start int64) int {
	var term = start
	var count int
	for {
		count++
		if term%2 == 0 {
			term /= 2
		} else {
			term = term*3 + 1
		}
		if term == 1 {
			count ++
			return count
		}
	}
}
