package util

import (
	"strconv"
	"strings"
)

func ToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	return int32(i), err
}

func ToInt64(s string) (uint64, error) {
	i, err := strconv.ParseUint(strings.TrimSpace(s), 0, 64)
	return uint64(i), err
}

func ToFloat64(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	return v, err
}
