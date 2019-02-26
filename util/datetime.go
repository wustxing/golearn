package util

import (
	"math"
	"strconv"
	"time"
)

func ParseTimestamp2String(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}

func GetWeekCountToday() int {
	t := GetCurrentTimestamp() + 8*3600
	t += 86400 * 3 // 1970-01-01 周四
	weekCount := math.Floor(float64(t / (86400 * 7)))
	return int(weekCount)
}

func GetCurrentTimestamp() int64 {
	return GetCurrentTime().Unix()
}

func GetCurrentMicroTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func GetCountDayNow() int32 {
	return GetCountDay(GetCurrentTimestamp())
}

func GetCountDay(seconds int64) int32 {
	count := (seconds + 8*3600) / (60 * 60 * 24)
	return int32(count)
}

func GetNowFormatHour() int32 {
	sHourNow := time.Now().Format("2006010215")
	hourNow, _ := strconv.Atoi(sHourNow)
	return int32(hourNow)
}
