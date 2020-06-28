package util

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
)

var crandmy = rand.New(cryptoSource{})

func CRand(n int) int {
	var src cryptoSource
	rnd := rand.New(src)
	return rnd.Intn(n)
}

func UUID() string {
	b := make([]byte, 16)
	_, err := crand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

type cryptoSource struct {
}

func (s cryptoSource) Seed(seed int64) {

}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
