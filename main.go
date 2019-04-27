package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type User struct {
	Name string `json:"name" field:"name"`
}

type Data struct {
	Data interface{}
}

func main() {
	u := &User{Name: "xujialong"}
	data, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	datamodel := &Data{
		Data: (*User)(nil),
	}
	typ := reflect.TypeOf(datamodel.Data)
	elem := typ.Elem()
	newUser := reflect.New(elem).Interface()
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(reflect.ValueOf(newUser))
	//v:=reflect.ValueOf(newUser).Elem()
	//for i := 0; i < v.NumField(); i++ {
	//	val := v.Field(i).Interface()
	//	fmt.Printf("%v\n",val)
	//}
	newUserTyp := reflect.TypeOf(newUser)
	fmt.Println(newUserTyp, elem)

	//value:=reflect.ValueOf(u)
	//elem:=typ.Elem()
	//
	//typeElem:=reflect.TypeOf(elem)
	//fmt.Println(elem.NumField(),typ,typ.Kind(),typ.Elem(),value,typeElem)
	sql, sql1 := model2SqlColumn(newUserTyp.Elem())
	fmt.Println(sql, sql1)

	//x:=(*User)(nil)
	//y:=(*Data)(nil)
	//fmt.Println(x==nil,x==y,y==nil)
}

func model2SqlColumn(m reflect.Type) (string, string) {
	//m := reflect.TypeOf(model)
	//m:=model
	result := ""
	valueNumString := ""
	if m != nil {
		slice := make([]string, 0, m.NumField())
		vSlice := make([]string, 0, m.NumField())
		for i := 0; i < m.NumField(); i++ {
			tag := m.Field(i).Tag.Get("field")
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

//func model2SqlColumnValue(m reflect.Value) (string, string) {
//	//m := reflect.TypeOf(model)
//	//m:=model
//	result := ""
//	valueNumString := ""
//	if m != nil {
//		slice := make([]string, 0, m.NumField())
//		vSlice := make([]string, 0, m.NumField())
//		for i := 0; i < m.NumField(); i++ {
//			tag := m.Field(i).Tag.Get("field")
//			if tag == "-" {
//				continue
//			}
//			slice = append(slice, tag)
//			vSlice = append(vSlice, "?")
//		}
//		result = strings.Join(slice, ", ")
//		valueNumString = strings.Join(vSlice, ",")
//	}
//	return result, valueNumString
//}
