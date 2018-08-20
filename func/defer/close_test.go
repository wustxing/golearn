package main

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

func Test_close(t *testing.T) {
	r := reader{}
	err := release(r)
	t.Log(err)

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

type reader struct {
}

func (r reader) Close() error {
	return errors.New("close error")
}

func release(r io.Closer) (err error) {
	defer func() {
		if err = r.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	return
}
