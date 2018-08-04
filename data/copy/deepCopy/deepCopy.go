package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"reflect"
)

type Test struct {
	Id       int
	Ptr      *int
	HelloMap map[int32]*int
}

func main() {
	c := 3
	T1 := &Test{
		Id:       1,
		Ptr:      &c,
		HelloMap: make(map[int32]*int),
	}
	T1.HelloMap[1] = &c
	v := reflect.ValueOf(T1).Elem()
	fmt.Println(reflect.ValueOf(T1), v.Type(), reflect.TypeOf(T1))
	copy := reflect.New(v.Type()).Interface()
	err := deepCopy(copy, T1)
	if err != nil {
		log.Fatal(err)
	}
	if T2, ok := copy.(*Test); ok {
		fmt.Println(T2)
	}
	fmt.Println(T1, copy)
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
