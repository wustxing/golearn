package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {
	str := `恭喜玩家<font color="#0067FF">%s</font>获得<font color="#ff4500">%s</font>X%s`
	msg := fmt.Sprintf(str, "xu", "jia", strconv.Itoa(10))
	fmt.Println(msg)

	x := int64(10000) / int64(344)
	fmt.Println(x)
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
