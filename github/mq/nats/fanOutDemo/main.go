package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"sync"
	"time"
)

const (
	workCnt = 50000
	chanLen = 1000000
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	msgChan := make(chan *nats.Msg, chanLen)
	nc.ChanSubscribe("log_topic", msgChan)

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(len(msgChan))
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(workCnt)
	go pool(&wg, workCnt, msgChan)
	wg.Wait()
}

func pool(wg *sync.WaitGroup, workerCnt int, tasksch <-chan *nats.Msg) {
	for i := 0; i < workerCnt; i++ {
		go worker(tasksch, wg)
	}
}

func worker(tasksch <-chan *nats.Msg, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, ok := <-tasksch
		if !ok {
			return
		}

		time.Sleep(time.Millisecond * 500)
	}
}
