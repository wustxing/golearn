package worker

import "sync"

type Worker struct {
	ch chan int
	wg sync.WaitGroup
}

func (p *Worker) Wait() {
	p.wg.Wait()
}

func (p *Worker) Add(f func()) {
	<-p.ch
	go func() {
		defer func() {
			p.ch <- 1
		}()
		f()
	}()
}

func NewWorker(chCount int) *Worker {
	p := &Worker{
		ch: make(chan int, chCount),
		wg: sync.WaitGroup{},
	}
	for i := 0; i < chCount; i++ {
		p.ch <- 1
	}
	return p
}
