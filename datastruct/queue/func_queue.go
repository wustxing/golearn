package queue

import (
	"fmt"
	"time"
)

type Locker interface {
	Lock(flag string)
	Unlock()
}

type locker struct {
	lock  bool
	timer *time.Timer
	queue *FuncQueue
	flag  string
}

func (l *locker) Lock(flag string) {
	l.flag = flag

	if l.lock {
		return
	}
	l.lock = true
	l.timer = time.AfterFunc(time.Second*5, func() {
		fmt.Printf("%s timeout %v\n", flag, time.Now().Unix())
		l.Unlock()
	})
}

func (l *locker) Unlock() {
	l.queue.Post(l.unlock)
}

func (l *locker) unlock() {
	if !l.lock {
		return
	}
	l.lock = false
	l.timer.Stop()
	l.queue.pop()
	l.queue.do()
}

type Worker interface {
	Post(f func())
	Len() int
}

type FuncQueue struct {
	funcs []func(locker Locker)
	Worker
}

func NewFuncQueue(worker Worker) *FuncQueue {
	return &FuncQueue{
		funcs:  nil,
		Worker: worker,
	}
}

func (q *FuncQueue) Push(f func(lk Locker)) {
	q.funcs = append(q.funcs, f)
	if len(q.funcs) == 1 {
		q.do()
	}
}

func (q *FuncQueue) do() {
	for len(q.funcs) > 0 {
		f := q.funcs[0]
		lk := &locker{
			lock:  false,
			timer: nil,
			queue: q,
		}

		f(lk)
		if lk.lock {
			break
		}
		q.pop()
	}
}

func (q *FuncQueue) pop() {
	copy(q.funcs, q.funcs[1:])
	q.funcs = q.funcs[:len(q.funcs)-1]
}
