package benchmark

import (
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
