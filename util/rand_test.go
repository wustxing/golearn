package util

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(8))
	fmt.Println(RandomStringRestrictions(8))
	fmt.Println(RandBetween(10, 100))
	fmt.Println(RandBetween('a', 'z'))
	fmt.Println(CRand(100))
	fmt.Println(UUID())
}

func TestShuffle(t *testing.T) {
	b := []byte{1, 2, 3, 4, 5}
	Shuffle(b)
	fmt.Println(b)
}

func BenchmarkCRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		crandmy.Intn(1000)
	}
}

func BenchmarkRank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(1000)
	}
}
