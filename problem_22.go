package main

import (
	"io/ioutil"
	"strings"
	"sort"
	"fmt"
)

var ci = map[rune]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'D': 4,
	'E': 5,
	'F': 6,
	'G': 7,
	'H': 8,
	'I': 9,
	'J': 10,
	'K': 11,
	'L': 12,
	'M': 13,
	'N': 14,
	'O': 15,
	'P': 16,
	'Q': 17,
	'R': 18,
	'S': 19,
	'T': 20,
	'U': 21,
	'V': 22,
	'W': 23,
	'X': 24,
	'Y': 25,
	'Z': 26,
}

func main() {
	f, err := ioutil.ReadFile("/home/hamed/dev/projects/euler_50_go/names.txt")
	if err != nil {
		panic(err)
	}

	names := strings.Split(string(f), ",")
	sort.Strings(names)
	sum := 0
	for i, name := range names {
		sum += worth(name) * (i + 1)
	}
	fmt.Println(sum)
}

func worth(name string) int {
	chars := []rune(name)
	sum := 0
	for _, c := range chars {
		if v, ok := ci[c]; ok {
			sum += v
		}
	}
	return sum
}
