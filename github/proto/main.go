package main

import (
	"fmt"
	"net/http"
)

var protoData = []byte{8,10,18,4,48,57,57,48,26,3,1,2,3,34,4,48,57,57,48,34,4,48,57,57,49}

func main(){
	http.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.RequestURI)
		fmt.Println(request.FormValue("cp_order_id"))
		fmt.Println(request.URL.Path)
	})
	http.ListenAndServe(":8989",nil)
}
