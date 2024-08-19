package common

import (
	"context"
	"encoding/json"
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

func GetNextMonthDate(start int64, n int) int64 {
	t := time.UnixMilli(start)

	nextMonth := t.AddDate(0, n, 0)

	lastDayOfNextMonth := time.Date(t.Year(), t.Month()+time.Month(n+1), 0, 0, 0, 0, 0, t.Location()).Day()
	if t.Day() > lastDayOfNextMonth {
		nextMonth = time.Date(t.Year(), t.Month()+time.Month(n), lastDayOfNextMonth, 0, 0, 0, 0, t.Location())
	}

	return nextMonth.UnixMilli()
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
