package main

import (
	"github.com/chenjiandongx/go-echarts/charts"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, _ *http.Request) {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-示例图"})
	m := make(map[string]interface{})
	m["工作进度"] = rand.Intn(50)
	gauge.Add("gauge", m)

	f, err := os.Create("gauge.html")
	if err != nil {
		log.Println(err)
	}
	gauge.Render(w, f)
}
