package worker

import (
	"fmt"
	"testing"
	"time"
)

func Test_Worker(t *testing.T) {
	t1 := time.Now()
	w := NewWorker(2)
	w.Add(func() {
		time.Sleep(time.Second * 1)
	})

	w.Add(func() {
		time.Sleep(time.Second * 1)
	})

	w.Add(func() {
		time.Sleep(time.Second * 1)
	})

	w.Wait()
	fmt.Println(time.Since(t1))
	return
}
