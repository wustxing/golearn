package zset

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestZSet(t *testing.T) {
	z := New()

	z.AddBatch(Element{
		Member: "a",
		Score:  100,
		Extra:  "aa",
	})

	z.AddBatch(Element{
		Member: "b",
		Score:  50,
		Extra:  "bb",
	})

	z.AddBatch(Element{
		Member: "c",
		Score:  200,
		Extra:  "cc",
	})

	z.AddBatch(Element{
		Member: "d",
		Score:  400,
		Extra:  "dd",
	})

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

	z.AddBatch(Element{
		Member: "c",
		Score:  900,
		Extra:  "cc",
	})

	if z.Rank("c", false) != 3 {
		t.Fail()
	}

	fmt.Println(z.V("c"))

	fmt.Println(z.Length())

	z.Incr("c", -20)
	score, _ = z.Score("c")
	if score != 880 {
		t.Fail()
	}

}

func Benchmark_ZSetAdd(b *testing.B) {
	z := New()
	for i := 0; i < b.N; i++ {
		z.Add(fmt.Sprintf("%d", i), i, "9999")
	}
}

func Benchmark_ZSetRank(b *testing.B) {
	z := New()
	for i := 0; i < 100000; i++ {
		z.Add(fmt.Sprintf("%d", i), i, "9999")
	}
	for i := 0; i < b.N; i++ {
		z.Rank("10000", false)
	}
}

func TestZSet_Rank(t *testing.T) {
	z := New()
	for i := 0; i < 100000; i++ {
		z.Add(fmt.Sprintf("%d", i), i, "9999")
	}
	for i := 200000; i > 0; i-- {
		z.Add(fmt.Sprintf("%d", i), i, "9999")
	}

	fmt.Println(z.Rank("200000", true))
}
