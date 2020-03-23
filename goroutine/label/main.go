package main

import (
	"context"
	"runtime/pprof"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {

		pprof.Do(context.Background(), pprof.Labels("helloxu", strconv.Itoa(i)), func(_ context.Context) {
			time.Sleep(time.Minute)
		})

	}
	time.Sleep(time.Hour)
}
