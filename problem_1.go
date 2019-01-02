package main

import "fmt"

func main() {
	var sum int64

	for i := int64(1000); i > 2; i-- {
		if i%3 == 0 || i%5 == 3 {
			sum += i
		}
	}

	fmt.Println("Sum: ", sum)
}
