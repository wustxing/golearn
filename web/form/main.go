package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func sayhellName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	fmt.Fprintf(w, "hello!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gptl")
		log.Println(t.Execute(w, token))
	} else {
		fmt.Println("username:", r.Form.Get("name")) //输出到服务器端
		fmt.Println("password:", r.Form.Get("password"))
		fmt.Println("token:", r.Form.Get("token"))
		//	template.HTMLEscape(w, []byte(r.Form.Get("name"))) //输出到客户端
		//fmt.Fprintf(w, r.Form.Get("name"))
		//fmt.Println("username", r.FormValue("name"))
		//fmt.Println("password", r.FormValue("password"))
		//t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		//err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	}
}

func main() {
	http.HandleFunc("/", sayhellName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
