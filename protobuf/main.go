package main

import (
	"github.com/0990/golearn/protobuf/pb"
	"fmt"
)

type Student struct{
	Name string
	Age int
}

//func write(){
//	contact := test.Contact{}
//
//	person1:=test.Person{
//		Id:1,
//		Name:"xu",
//		PhoneType:test.PhoneType_HOME,
//	}
//
//	person2 :=test.Person{
//		Id:2,
//		Name:"jia",
//		PhoneType:test.PhoneType_WORK,
//	}
//	//fmt.Println(person1)
//
//	contact.Person = append(contact.Person, &person1)
//	contact.Person = append(contact.Person,&person2)
//
//	data,_ := msg.Marshal(&contact)
//
//	str:=msg.MarshalTextString(&contact)
//
//	conta:=test.Contact{}
//	msg.UnmarshalText(str,&conta)
//
//	fmt.Println(str,conta)
//
//
//	ioutil.WriteFile("./test.txt",data,os.ModePerm)
//}
//
//func read(){
//	data,err := ioutil.ReadFile("./test.txt")
//
//	if err!=nil{
//			log.Fatalln(err)
//	}
//
//	contact:=test.Contact{}
//	msg.Unmarshal(data,&contact)
//
//	for _,v := range contact.Person{
//		fmt.Println(v)
//	}
//	fmt.Printf("%+v",contact)
//}

func main(){
	value:=test.MapValue{}
	value.Name = "xu"
	value.Age = 15
	//err:=msg.UnmarshalText("",&value)
	//fmt.Println(err,value)
	fmt.Println(value)
	//write()
	//read()
}