package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 2, 50)
	wg.Wait()
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)
	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}
	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}
	close(tasksCh)
}

func worker(tasksch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-tasksch
		if !ok {
			return
		}

		//d := time.Duration(task) * time.Millisecond
		time.Sleep(time.Second)
		fmt.Println("processing task", task)
	}
}
