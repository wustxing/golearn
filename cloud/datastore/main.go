package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"log"
	"net/http"
	"os"
)

type Entity struct {
	Value string
}

func main() {
	http.HandleFunc("/", handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "enter handle\n")
	ctx := appengine.NewContext(r)
	fmt.Fprintf(w, "create ctx\n")
	k := datastore.NewKey(ctx, "Entity", "stringID", 0, nil)

	e := new(Entity)

	if err := datastore.Get(ctx, k, e); err != nil {
		fmt.Fprintf(w, "datastore get error:%v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}
	old := e.Value
	e.Value = r.URL.Path
	if _, err := datastore.Put(ctx, k, e); err != nil {
		fmt.Fprintf(w, "datastore put error:%v\n", err)
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "text/plain;charset=utf-8")
	fmt.Fprintf(w, "old=%q\nnew=%q\n", old, e.Value)
}
