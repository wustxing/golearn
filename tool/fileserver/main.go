package main

import (
	"flag"
	"fmt"
	"github.com/0990/golearn/util"
	auth "github.com/abbot/go-http-auth"
	"net/http"
)

var port = flag.Int("port",8080,"listen port")

func secret(user,realm string)string{
	if user=="0990"{
		return "e81053232e165aee62cfc3aa948109d4"
	}
	return ""
}

func handle(w http.ResponseWriter,r *auth.AuthenticatedRequest){
	fmt.Fprintf(w,"hi",r.Username)
}

func main(){
	flag.Parse()
	listenAddr:=fmt.Sprintf("0.0.0.0:%d",*port)

	authenticator:=auth.NewDigestAuthenticator("example.com",secret)
	http.Handle("/",authenticator.JustCheck(http.FileServer(http.Dir("")).ServeHTTP))


	err:=http.ListenAndServe(listenAddr,nil)
	fmt.Println(err)
}

func makeSecret(userName string,realm string,password string)string{
	str:=userName+":"+realm+":"+password
	return util.MD5(str)
}