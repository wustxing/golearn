package main

import (
	"fmt"
	"strconv"
	"time"
)

type Receipt struct {
	BundleId string      `json:"bundle_id"`
	InApp    []*InAppItem `json:"in_app"`
}

type InAppItem struct{
	Id int32
}

func main(){
	fmt.Println(time.Now().UnixNano())
}

func GetDayIndexAfter2000ByUnix(sec int64) int32 { //180101 18年1月1日
	if sec == 0 {
		return 0
	}
	t := time.Unix(sec, 0)
	sDayNow := t.Format("20060102")
	dayNow, _ := strconv.Atoi(sDayNow)
	if dayNow > 20000000 {
		return int32(dayNow - 20000000)
	}
	return 0
}

func ParseDayIndexAfter2000(day int32)(time.Time,error){
	dayIndex:=day+20000000
	loc, _ := time.LoadLocation("Local")
	t, err := time.ParseInLocation("20060102", strconv.FormatInt(int64(dayIndex), 10), loc)
	if err != nil {
		return time.Time{}, err
	}
	return t,nil
}