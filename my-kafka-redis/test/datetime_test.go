package test

import (
	"testing"
	"time"
	"fmt"
)

func Test_DateTime(t *testing.T) {

	duration := time.Duration(3 * 3600 * 24 * 1000000000)
	t.Log(duration.String())
	t.Logf("%f - %f - %f - %d", duration.Hours(), duration.Minutes(), duration.Seconds(), duration.Nanoseconds())
	t.Log("-------------------------------")
	now := time.Now()
	hourly := time.Unix(now.Unix(), 0).UTC().Format("2006-01-02T15:00:00.000Z")

	loc, _ := time.LoadLocation("Asia/Shanghai")
	curDay := now.Format("2006-01-02 00:00:00")
	dayTime, _ := time.ParseInLocation("2006-01-02 15:04:05", curDay, loc)
	daily := dayTime.UTC().Format("2006-01-02T15:04:05.000Z")

	monthTime := fmt.Sprintf("%d-0%d-0%d 00:00:00", now.Year(), now.Month(), 01)
	t.Log(monthTime)
	month, _ := time.ParseInLocation("2006-01-02 15:04:05", monthTime, loc)
	monthly := month.UTC().Format("2006-01-02T15:04:05.000Z")
	t.Log(hourly)
	t.Log(daily)
	t.Log(monthly)
}
