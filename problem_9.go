package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 1.0; i <= 10000; i++ {
		for j := 1.0; j <= 10000; j++ {
			sum := i + j + math.Sqrt(float64((i*i)+(j*j)))
			if sum == 1000 {
				fmt.Printf("%f", i*j*math.Sqrt(float64((i*i)+(j*j))))
				return
			}
		}
	}
}
