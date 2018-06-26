package main

import (
	"fmt"
	"github.com/0990/golearn/tabtoy/table"
)

func main(){
	config:=table.NewConfigTable()


	if err:=config.Load("Config.json");err!=nil{
		panic(err)
	}

	for index,v:=range config.Sample {
		fmt.Println(index,v,v.Name)
		fmt.Println(v.Pos.X)
		//fmt.Println(v.Name,v.Age,v.ID)
	}
	fmt.Println(config.Sample[0].Pos.X);
	fmt.Println(config.Sample)

	//data,err := ioutil.ReadFile("Config.pbt")
	//
	//if err!=nil{
	//	panic(err)
	//}
	//
	//pbConfig := &table.Config{}
	//err = msg.UnmarshalText(string(data),pbConfig)
	//
	//fmt.Println(pbConfig,err)
	//
	//for index,v := range pbConfig.GetSample(){
	//	fmt.Println(index,v.ID,v.Age,v.Name)
	//}
}
