package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Prop struct {
	propID int32
	num    int32
}

func main() {
	fields := make(map[string]string)

	fields["hello"] = "m"
	fields["world"] = "j"

	fieldstring := make([]string, 0, len(fields))

	for k, _ := range fields {
		fieldstring = append(fieldstring, k)
	}

	fmt.Println(strings.Join(fieldstring, ",,"))

	fmt.Println(ParseProp("2, 1 ;5 ,4 "))
}

func ToInt32(s string) (int32, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))
	return int32(i), err
}

func ParseProp(s string) (propList []Prop) {
	fields := strings.Split(s, ";")
	for _, v := range fields {
		propSlice := strings.Split(v, ",")
		if len(propSlice) == 2 {
			propID, perr := ToInt32(propSlice[0])
			num, nerr := ToInt32(propSlice[1])
			if perr == nil && nerr == nil {
				propList = append(propList, Prop{
					propID: propID,
					num:    num,
				})
			}
		}
	}
	return
}
