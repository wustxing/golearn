package queue

import (
	"fmt"
	"github.com/0990/golearn/rpc/service"
	"sync"
	"testing"
	"time"
)

func TestFuncQueue(t *testing.T) {
	w := service.NewWorker()
	w.Run()

	q := NewFuncQueue(w)

	var wg sync.WaitGroup
	wg.Add(1)
	var i int

	w.Post(func() {
		q.Push(func(lk Locker) {
			i++
			lk.Lock("1")
			time.AfterFunc(3*time.Second, func() {
				w.Post(func() {
					lk.Unlock()
					if i != 1 {
						t.Fail()
					}
				})
			})
		})

		q.Push(func(lk Locker) {

			i++
			lk.Lock("2")
			time.AfterFunc(time.Second, func() {
				w.Post(func() {

					lk.Unlock()
					if i != 2 {
						t.Fail()
					}
				})
			})
		})
	})

	w.Post(func() {

		q.Push(func(lk Locker) {

			i++
			lk.Lock("3")
			time.AfterFunc(time.Second, func() {
				w.Post(func() {
					defer wg.Done()
					lk.Unlock()
					if i != 3 {
						t.Fail()
					}
				})
			})
		})
	})

	wg.Wait()
}

func TestFuncQueue_Elapse(t *testing.T) {
	w := service.NewWorker()
	w.Run()

	q := NewFuncQueue(w)
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		w.Post(func() {
			q.Push(func(lk Locker) {
				lk.Lock("hello")
				w.Post(func() {
					lk.Unlock()
					wg.Done()
				})
			})
		})
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}
