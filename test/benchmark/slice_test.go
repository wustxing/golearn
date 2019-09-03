package benchmark

import "testing"

const size = 100000 //See size values benchmarked later in table
func BenchmarkEmptyInit(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkKnownAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, 0, size)
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
}

func BenchmarkKnownDirectAccess(b *testing.B) {
	for n := 0; n < b.N; n++ {
		data := make([]int, size, size)
		for k := 0; k < size; k++ {
			data[k] = k
		}
	}
}
