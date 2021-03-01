package main

import (
	"fmt"
	auth "github.com/abbot/go-http-auth"
	"net/http"
)

func Secret(user,realm string)string{
	if user=="john"{
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
	}
	return ""
}

func handle(w http.ResponseWriter,r *auth.AuthenticatedRequest){
	fmt.Fprintf(w,"hello",r.Username)
}

func main(){
	authenticator:=auth.NewBasicAuthenticator("",Secret)

	http.HandleFunc("/",authenticator.Wrap(handle))
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		panic(err)
	}
}
