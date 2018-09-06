package main

import (
	"fmt"
	"math"
	"time"
)

const (
	goBornDate       = "2006-01-02"          //长日期格式
	goBornShortDate  = "06-01-02"            //短日期格式
	goBornTimes      = "15:04:05"            //长时间格式
	goBornShortTime  = "15:04"               //短时间格式
	goBornDateTime   = "2006-01-02 15:04:05" //日期时间格式
	goBornDateString = "20060102"
)

func main() {
	fmt.Println(GetCountDay())
	year, month, _ := time.Now().Date()
	targetDayEndTime := time.Date(year, month, 3, 23, 59, 59, 0, time.Local)
	fmt.Println(targetDayEndTime)
	fmt.Println(time.Now().Day())
	fmt.Println(time.Now(), time.Now().Unix())
	fmt.Println(time.Now().UTC(), time.Now().UTC().Unix())
	fmt.Println(GetTodayZeroTime(), GetTodayZeroTime().Unix())
	fmt.Println(GetTodayZeroTime().UTC(), GetTodayZeroTime().Unix())
	fmt.Println(time.Now().Unix() - GetTodayZeroTime().Unix())
}

func GetTodayZeroTime() time.Time {
	timeStr := GetCurrentTime().Format(goBornDate)
	zeroTime, _ := time.ParseInLocation(goBornDate, timeStr, time.Local)
	return zeroTime
}

func GetCountDay() int32 {
	t := GetCurrentTimestamp()
	count := math.Floor(float64(t / (60 * 60 * 24)))
	return int32(count)
}

func GetCurrentTimestamp() int64 {
	return GetCurrentTime().Unix()
}

func GetCurrentTime() time.Time {
	return time.Now()
}
