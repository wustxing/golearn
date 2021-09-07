package skiplist

import "strconv"

func toString(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func toStringInt64(v uint64) string {
	return strconv.FormatInt(int64(v), 10)
}
