package main

import (
	"testing"
	"time"
)

func Test_Main(t *testing.T) {
	go func() {
		main()
	}()
	time.Sleep(time.Hour)
}
