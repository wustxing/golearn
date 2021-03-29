package main

import (
	"bytes"
	"context"
	"encoding/json"
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://127.0.0.1:9200",
		},
	}
	es,err:=elasticsearch.NewClient(cfg)
	if err!=nil{
		log.Fatal(err)
	}

	res,err:=es.Info()
	if err!=nil{
		log.Fatal("Error getting response: %s",err)
	}

	defer res.Body.Close()

	var b strings.Builder
	b.WriteString(`{"m.level":"`)
	b.WriteString("error")
	b.WriteString(`"}`)

	req:=esapi.IndexRequest{
		Index:"2020.11.10",
		DocumentID: "",
		Body:strings.NewReader(b.String()),
		Refresh: "true",
	}

	res,err=req.Do(context.Background(),es)
	if err!=nil{
		log.Fatalf("Error getting response:%s",err)
	}
	defer res.Body.Close()

	if res.IsError(){
		log.Printf("[%s] Error indexing document ID=%d",res.Status(),1)
	}else{
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	var buf bytes.Buffer
	query:=map[string]interface{}{
		"query":map[string]interface{}{
			"match":map[string]interface{}{
				"m.level": "error",
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res,err=es.Search(es.Search.WithContext(context.Background()),
		es.Search.WithIndex("2020.11.10"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty())
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var (
		r  map[string]interface{}
	)


	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}

