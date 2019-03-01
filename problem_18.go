package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)


// this one uses a greedy algorithm selecting the best option
// at each step regardless of anything other steps ahead
// isn't the correct solution though
func main() {
	f, err := os.Open("problem_18.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	intsLin := make([]int, 0)
	for l := range lines {
		digits := strings.Split(lines[l], " ")
		for _, c := range digits {
			i, _ := strconv.Atoi(c)
			intsLin = append(intsLin, i)
		}
	}

	//fmt.Println(intsLin)

	sum := 0
	selectedIndex := 0
	for l := range lines {
		fmt.Println(intsLin[selectedIndex])
		sum += intsLin[selectedIndex]

		nextIndex := selectedIndex + l + 1
		if nextIndex > len(intsLin) {
			break
		}
		if intsLin[nextIndex] > intsLin[nextIndex+1] {
			selectedIndex = nextIndex
		} else {
			selectedIndex = nextIndex + 1
		}
	}

	fmt.Println(sum)

}
