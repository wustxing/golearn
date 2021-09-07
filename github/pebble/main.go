package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/cockroachdb/pebble"
	"sync"
	"time"
)

func main() {
	db, err := pebble.Open("test", &pebble.Options{})
	if err != nil {
		panic(err)
	}

	data := []byte("hello")

	concurrent := 1

	var wg sync.WaitGroup
	wg.Add(concurrent)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	counts := make([]int, concurrent)
	start := time.Now()
	for j := 0; j < concurrent; j++ {
		index := uint64(j)
		go func() {
			count := 0
			i := index
		LOOP:
			for {
				select {
				case <-ctx.Done():
					break LOOP
				default:
					db.Set(genKey(i), data, &pebble.WriteOptions{})
					i += uint64(concurrent)
					count++
				}
			}
			counts[index] = count
			wg.Done()
		}()
	}
	wg.Wait()
	dur := time.Since(start)
	d := int64(dur)
	var n int
	for _, count := range counts {
		n += count
	}
	fmt.Printf("set rate: %d op/s, mean: %d ns, took: %d s\n", int64(n)*1e6/(d/1e3), d/int64((n)*(concurrent)), int(dur.Seconds()))
}

func genKey(i uint64) []byte {
	r := make([]byte, 9)
	r[0] = 'k'
	binary.BigEndian.PutUint64(r[1:], i)
	return r
}
