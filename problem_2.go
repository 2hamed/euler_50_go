package main

import "fmt"

func main() {
	var a int64 = 0
	var b int64 = 1
	var sum int64
	for {
		t := a + b
		a = b
		b = t
		if t%2 == 0 {
			sum += t
		}
		if t > 4000000 {
			break
		}
	}
	fmt.Println("Sum: ", sum)
}