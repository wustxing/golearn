package main

import (
	"fmt"
	"github.com/0990/golearn/util"
	"math"
	"time"
)

func main() {
	go func() {
		fmt.Println(util.GetGID())
	}()
	fmt.Println(util.GetGID())
	time.Sleep(time.Second)
}

func GetWeekCountToday() int {
	t := time.Now().UTC().Unix()
	t += 86400 * 3 // 1970-01-01 周四
	weekCount := math.Floor(float64(t / (86400 * 7)))
	return int(weekCount)
}

func GetTargetDayLastWeekStartZeroTime(t time.Time) time.Time {
	weekDay := GetWeekDay(t)
	year, month, day := t.Date()
	thisWeek := time.Date(year, month, day-weekDay, 0, 0, 0, 0, time.Local)
	lastWeek := thisWeek.AddDate(0, 0, -7)
	return lastWeek
}

// GetWeekDay 获取指定时间所属星期.
func GetWeekDay(t time.Time) int {
	return int(t.Weekday())
}
