package zset

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestZSet(t *testing.T) {
	z := New()

	z.Add("a", 100)
	z.Add("b", 50)
	z.Add("c", 200)
	z.Add("d", 400)

	ret := z.Range(1, -2, false)
	data, err := json.Marshal(ret)
	fmt.Printf("%v %v", string(data), err)

	if z.Rank("c", false) != 2 {
		t.Fail()
	}

	if z.Rank("c", true) != 1 {
		t.Fail()
	}

	score, _ := z.Score("d")
	if score != 400 {
		t.Fail()
	}

	z.Add("c", 900)

	if z.Rank("c", false) != 3 {
		t.Fail()
	}

	z.Incr("c", -20)
	score, _ = z.Score("c")
	if score != 880 {
		t.Fail()
	}
}

func Benchmark_ZSetAdd(b *testing.B) {
	z := New()
	for i := 0; i < b.N; i++ {
		z.Add(fmt.Sprintf("%d", i), int64(i))
	}
}

func Benchmark_ZSetRank(b *testing.B) {
	z := New()
	for i := 0; i < 100000; i++ {
		z.Add(fmt.Sprintf("%d", i), int64(i))
	}
	for i := 0; i < b.N; i++ {
		z.Rank("10000", false)
	}
}

func TestZSet_Rank(t *testing.T) {
	z := New()
	for i := 0; i < 100000; i++ {
		z.Add(fmt.Sprintf("%d", i), int64(i))
	}
	for i := 200000; i > 0; i-- {
		z.Add(fmt.Sprintf("%d", i), int64(i))
	}

	fmt.Println(z.Rank("200000", true))
}
