package common

import (
	"context"
	"encoding/json"
	"math"
	"time"
)

func GetUserIDFromContext(ctx context.Context) (userID int64, err error) {
	ret, err := ctx.Value(("userID")).(json.Number).Int64()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

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

func GetBillIndexByTime(start, current int64) int64 {
	if start >= current {
		return 0
	}

	// returns the number of months difference
	t1 := time.UnixMilli(start)
	t2 := time.UnixMilli(current)

	y1, m1, _ := t1.Date()
	y2, m2, _ := t2.Date()

	months := (y2-y1)*12 + int(m2) - int(m1)

	return int64(months)
}
