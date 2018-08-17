package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

const (
	workCnt = 100
	chanLen = 1000000000
)

var totalCnt int32

type consumer struct {
	msgs []*nats.Msg
}

func NewConsumer() *consumer {
	return &consumer{
		msgs: make([]*nats.Msg, 0),
	}
}

func (p *consumer) consume(tasksch <-chan *nats.Msg) {
	ticker := time.NewTicker(time.Second * 5)
	defer func() {
		fmt.Println("consumer leave")
		ticker.Stop()
		wg.Done()
	}()
	for {
		select {
		case msg, ok := <-tasksch:
			if !ok {
				p.saveTodb()
				return
			}
			p.msgs = append(p.msgs, msg)
			if len(p.msgs) > 1000 {
				p.saveTodb()
			}
		case <-ticker.C:
			p.saveTodb()
		}
	}
}

func (p *consumer) saveTodb() {
	if len(p.msgs) <= 0 {
		return
	}
	atomic.AddInt32(&totalCnt, int32(len(p.msgs)))
	p.msgs = []*nats.Msg{}
	time.Sleep(time.Millisecond * 1000)
}

var wg sync.WaitGroup

func main() {

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	msgChan := make(chan *nats.Msg, chanLen)
	sub, _ := nc.ChanSubscribe("log_topic", msgChan)

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println(len(msgChan))
			}
		}
	}()
	s := <-exit
	fmt.Println("Got signal:", s, "chanLen:", len(msgChan))
	wg.Add(workCnt)
	go pool(workCnt, msgChan)
	sub.Unsubscribe()
	close(msgChan)
	wg.Wait()
	fmt.Println("exit,ReceiveTotalLen:", totalCnt)
}

func pool(workerCnt int, tasksch <-chan *nats.Msg) {
	for i := 0; i < workerCnt; i++ {
		consumer := NewConsumer()
		go consumer.consume(tasksch)
	}
}

//func worker(tasksch <-chan *nats.Msg, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	for {
//		_, ok := <-tasksch
//		if !ok {
//			return
//		}
//
//		time.Sleep(time.Millisecond * 500)
//	}
//}
