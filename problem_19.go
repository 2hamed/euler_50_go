package main

import "fmt"

func main() {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	months_leap := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	week_day := 1
	sundays_on_first_day := 0

	for y := 1900; y <= 2000; y++ {
		var m []int
		if y%4 == 0 && y%100 != 0 {
			m = months_leap
		} else if y%400 == 0 && y%100 == 0 {
			m = months_leap
		} else {
			m = months
		}
		for _, days := range m {
			for d := 1; d <= days; d++ {
				if d == 1 && week_day == 7 && y > 1900 {
					sundays_on_first_day++
				}
				week_day++
				if week_day > 7 {
					week_day = 1
				}
			}
		}
	}

	fmt.Println(sundays_on_first_day)
}
