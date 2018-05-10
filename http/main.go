package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string(s))
}

type Struct struct {
	Greeting string
	Punch    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	formatValue := fmt.Sprintf("%v,%v,%v", s.Greeting, s.Punch, s.Who)
	fmt.Fprint(w, formatValue)
}

func main() {
	http.Handle("/", String("I am xujialong"))
	http.Handle("/struct", &Struct{"Hello", ":", "Xujialong"})

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		log.Fatal(err)
	}

}
