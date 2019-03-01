package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"fmt"
)

// this algorithm approaches the matrix from bottom to top adding
// the largest child to it's parent coming up. the solution will at the very top
func main() {
	f, err := os.Open("problem_67.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	max := len(lines)
	matrix := make([][]int, max)
	for l := range lines {
		digits := strings.Split(lines[l], " ")
		matrix[l] = make([]int, max)
		for j, c := range digits {
			i, _ := strconv.Atoi(c)
			matrix[l][j] = i
		}
	}

	for i := len(matrix) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if matrix[i][j] > matrix[i][j+1] {
				matrix[i-1][j] += matrix[i][j]
			} else {
				matrix[i-1][j] += matrix[i][j+1]
			}
		}
	}

	fmt.Println(matrix[0][0])

}
