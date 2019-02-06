package main

import (
	"math"
	"fmt"
	"strconv"
)

func main() {
	r := math.Pow(2, 1000)
	str := fmt.Sprintf("%.0f\n", r)
	sum := 0
	for _, c := range str {
		n, _ := strconv.Atoi(string(c))
		sum += n
	}

	fmt.Println(sum)

}
