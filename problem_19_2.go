package main

import (
	"time"
	"fmt"
)

func main() {
	date := time.Date(1901, 1, 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	sundays_on_first_day := 0
	for date.Before(date2) {
		if date.Weekday() == time.Sunday && date.Day() == 1 {
			sundays_on_first_day++
		}
		date = date.Add(24 * time.Hour)
	}

	fmt.Println(sundays_on_first_day)
}
