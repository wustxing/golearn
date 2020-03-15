package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/sohlich/elogrus.v2"
	"math/rand"
	"time"

	"github.com/davyxu/cellnet"
)

func initLog() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewElasticHook(client, "localhost", log.DebugLevel, "mylog")
	if err != nil {
		log.Panic(err)
	}
	log.AddHook(hook)
}

func main() {
	cellnet.NewEventQueue()

	rand.Seed(time.Now().UnixNano())
	log.SetFormatter(new(log.JSONFormatter))
	initLog()
	for {
		r := rand.Int31n(100)
		if r < 50 {
			log.WithField("me", r).Error(time.Now().Unix())
		} else {
			log.WithField("mr", r).Info(time.Now().Unix())
		}
		time.Sleep(5 * time.Second)
	}

	time.Sleep(5 * time.Second)
}

//http://localhost:9200/mylog/_search
