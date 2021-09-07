package main

import (
	"context"
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/0990/golearn/datastruct/skiplist"
	"sync"
	"time"
)

var concurrent_cnt = flag.Int("c", 2, "concurrent count")
var storeType = flag.String("store", "memory", "storeType")

func main() {
	flag.Parse()
	store, err := skiplist.NewKVStore(*storeType, "test")
	if err != nil {
		panic(err)
	}

	defer store.Close()

	concurrent := *concurrent_cnt

	var wg sync.WaitGroup
	wg.Add(concurrent)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	counts := make([]int, concurrent)
	start := time.Now()
	for j := 0; j < concurrent; j++ {
		index := uint64(j)

		z1 := skiplist.NewZSet(store, fmt.Sprintf("%v", index))
		go func() {
			count := 0
			i := index
		LOOP:
			for {
				select {
				case <-ctx.Done():
					break LOOP
				default:
					z1.ZAdd(string(genKey(i)), float64(i))
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
