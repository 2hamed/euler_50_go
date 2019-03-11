package main

import "fmt"

func main() {
	fmt.Println(string(revers([]rune("1234567890"))))
}

func revers(r []rune) []rune {
	if len(r) == 1 {
		return append(make([]rune, 0), r[0])
	}

	return append(revers(r[1:]), r[0])
}
