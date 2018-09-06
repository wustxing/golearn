package main

import (
	"net/http"
	"fmt"
)

func main(){

	http.HandleFunc("/",func(w http.ResponseWriter,r*http.Request){
		fmt.Fprint(w,"hello docker")
	})
	fmt.Println("http listen start...")
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		fmt.Print(err)
	}
}
