package main

import (
	"fmt"
	"github.com/0990/golearn/util"
	auth "github.com/abbot/go-http-auth"
	"net/http"
)

func secret(user,realm string)string{
	if user=="john"{
		return "b98e16cbc3d01734b264adba7baa3bf9"
	}
	return ""
}

func handle(w http.ResponseWriter,r *auth.AuthenticatedRequest){
	fmt.Fprintf(w,"hi",r.Username)
}

func main(){
	str:="0990"+":"+"example.com"+":"+"110112"
	fmt.Println(auth.H(str))
	fmt.Println(util.MD5(str))
	auther:=auth.NewDigestAuthenticator("example.com",secret)
	http.Handle("/",auther.Wrap(handle))
	err:=http.ListenAndServe(":8800",nil)
	fmt.Println(err)
}
