package main

import "fmt"

// this solution's bound to find the it eventually but is way too slow

func main() {
	tNumber := 0
	for i := 1; true; i++ {
		tNumber += i
		f := factorsCount(int64(tNumber))
		fmt.Printf("%d: %d -> %d \n", i, tNumber, f)
		if f > 500 {
			fmt.Println("Number: ", tNumber)
			return
		}
	}
}

func factorsCount(n int64) int {
	count := 0
	for i := n; i > 0; i-- {
		if n%i == 0 {
			count++
		}
	}
	return count
}
