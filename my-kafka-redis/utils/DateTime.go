package utils

import (
	"time"
	"fmt"
)

const LongFormat = "2006-01-02T15:04:05.000Z"
const DefaultFormat = "2006-01-02T00:00:00Z"
const HourlyFormat = "2006-01-02T15:00:00.000Z"

var Local, _ = time.LoadLocation("Asia/Shanghai")

//parse Unix-Time to Golang Time
func UnixToTime(timestamp int64) time.Time {
	return time.Unix(timestamp, 0).UTC()
}

func UnixToTimeString(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).UTC().Format(format)
}

func Hourly(timestamp int64) string {
	return UnixToTimeString(timestamp, HourlyFormat)
}

func Daily(timestamp int64) string {
	dayTime := UnixToTime(timestamp).Format("2006-01-02 00:00:00")
	day, _ := time.ParseInLocation("2006-01-02 15:04:05", dayTime, Local)
	daily := day.UTC().Format(LongFormat)
	return daily
}

func Monthly(timestamp int64) string {
	date := UnixToTime(timestamp)
	monthTime := fmt.Sprintf("%d-0%d-0%d 00:00:00", date.Year(), date.Month(), 01)
	month, _ := time.ParseInLocation("2006-01-02 15:04:05", monthTime, Local)
	monthly := month.UTC().Format(LongFormat)
	return monthly
}
