package main

import (
	"fmt"
	"strconv"
	"strings"
)

func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func main() {
	s := []byte(`{"HelloWorld": "\uB155, \uC138\uC0C1(\u4E16\u4E0A). \u263a"}`)
	v, _ := UnescapeUnicode(s)
	fmt.Println(string(s), string(v))
}
