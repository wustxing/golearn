package benchmark

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

//go test -v -bench=Alloc -benchmem test/brenchmark/benchmark_test.go
// -bench指定只测试Alloc
func Benchmark_Alloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			m := j
			m++
			//只有分配到栈上的变量才算Alloc
			go func() {
				m++
			}()
		}
	}
}

func Benchmark_FileRename(b *testing.B) {
	f, err := os.OpenFile("text", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		b.Fatal(err)
	}
	f.Close()

	old := "text"
	for i := 0; i < b.N; i++ {
		new := strconv.FormatInt(int64(i), 10)
		err = os.Rename(old, new)
		if err != nil {
			fmt.Println(err)
		}
		old = new
	}
}

func Benchmark_FileWrite(b *testing.B) {
	f, err := os.OpenFile("test.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if i >= (b.N)-1 {
			WriteOffset(f, int64(i))
		}
	}
}

func WriteOffset(f *os.File, offset int64) error {
	err := f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}
	_, err = f.WriteString(strconv.FormatInt(offset, 10))
	if err != nil {
		return err
	}
	return nil
}
