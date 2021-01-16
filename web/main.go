package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	writer := NewLineBufferedWriter(os.Stdout)
	defer writer.Flush()

	done := make(chan int)

	go func() {
		fmt.Fprint(writer, "你好")
		done <- 1
	}()
	m := <-done

	println(m)
}

type LineBufferedWriter struct {
	*bufio.Writer
}

func NewLineBufferedWriter(w io.Writer) *LineBufferedWriter {
	return &LineBufferedWriter{
		Writer: bufio.NewWriter(w),
	}
}

func (w *LineBufferedWriter) Write(p []byte) (n int, err error) {
	n, err = w.Writer.Write(p)
	if err != nil {
		return n, err
	}

	if bytes.Contains(p, []byte{'\n'}) {
		w.Flush()
	}

	return n, err
}