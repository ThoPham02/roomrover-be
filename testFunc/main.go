package main

import (
	"fmt"
	"math"
	"time"
)

func GetCurrentTime() int64 {
	return time.Now().UnixMilli()
}

func GetNextMonthDate(start int64) int64 {
	var currentTime = GetCurrentTime()
	var n int = int(math.Round(float64(currentTime-start) / float64(86400000*30))) + 1

	t := time.UnixMilli(start)
	year, month, day := t.Date()
	if (int(month) + n) > 12 {
		year += (int(month) + n) / 12
		month = time.Month((int(month) + n) % 12)
	} else {
		month += time.Month(n)
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day > 31 {
			day = 31
		}
	case 4, 6, 9, 11:
		if day > 30 {
			day = 30
		}
	case 2:
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			if day > 29 {
				day = 29
			}
		} else {
			if day > 28 {
				day = 28
			}
		}
	}
	nextMonth := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).UnixMilli()
	return nextMonth
}

func main() {
	fmt.Println(GetNextMonthDate(1728820800000))
}
