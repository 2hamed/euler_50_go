package main

import (
	"fmt"
	"strconv"
	"math/big"
)

func main() {
	r := fact(100).String()
	runes := []rune(r)
	fmt.Println(sum(runes))
}

func fact(n int) *big.Int {
	if n == 1 {
		return big.NewInt(int64(1))
	}
	r := new(big.Int)
	r.Mul(big.NewInt(int64(n)), fact(n-1))
	return r
}
func sum(r []rune) int {
	var p = 0
	for _, v := range r {
		n, _ := strconv.Atoi(string(v))
		p += n
	}
	return p
}
