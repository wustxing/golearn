package main

import (
	"fmt"
	"reflect"
	"strings"
)

type modelLog struct {
	AreaID int32  `boil:"areaID" json:"id" toml:"id" yaml:"id"`
	Name   string `boil:"name" json:"id" toml:"id" yaml:"id"`
}

func main() {
	log := modelLog{
		AreaID: 1,
		Name:   "xu",
	}
	fmt.Println(model2SqlColumn(log))
}

func model2SqlColumn(model interface{}) (string, string) {
	m := reflect.TypeOf(model)
	result := ""
	valueNumString := ""
	if m != nil {
		slice := make([]string, 0, m.NumField())
		vSlice := make([]string, 0, m.NumField())
		for i := 0; i < m.NumField(); i++ {
			tag := m.Field(i).Tag.Get("boil")
			if tag == "-" {
				continue
			}
			slice = append(slice, tag)
			vSlice = append(vSlice, "?")
		}
		result = strings.Join(slice, ", ")
		valueNumString = strings.Join(vSlice, ",")
	}
	return result, valueNumString
}
