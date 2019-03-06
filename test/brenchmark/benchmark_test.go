package brenchmark

import (
	"fmt"
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
		fmt.Sprintf("%d", i)
	}
}
