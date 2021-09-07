package skiplist

import "testing"

func TestFloat64(t *testing.T) {
	a := 456.789
	data := float64Encode(a)
	b, ok := float64Decode(data)
	if !ok {
		t.FailNow()
	}
	if a != b {
		t.Fatal("a!=b")
	}
}
