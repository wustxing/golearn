package util

import "time"

func ParseTimestamp2String(sec int64) string {
	return time.Unix(sec, 0).Format("2006-01-02 15:04:05")
}
