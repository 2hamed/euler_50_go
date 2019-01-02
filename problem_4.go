package main

import (
	"strconv"
	"fmt"
	"time"
)

func main() {
	a := 999
	max := -1
	start := time.Now().UnixNano()

	for j := a; j >= 100; j-- {
		for i := a; i >= 100; i-- {
			p := j * i
			if isPalindrome2(p) {
				if p > max {
					max = p
				}
			}
		}
	}
	fmt.Println(max)
	fmt.Println("time=", time.Now().UnixNano()-start)
}
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	r := Reverse(s)
	//fmt.Println(s,"==",r)
	return s == r
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func isPalindrome2(n int) bool {
	a := make([]int, 0)
	b := n
	for b > 0 {
		a = append(a, b%10)
		b = b / 10
	}
	//fmt.Println(a)
	for i := 0; i < len(a)/2; i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}
	return true
}
