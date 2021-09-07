package skiplist

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/cockroachdb/pebble"
	"sync"
	"testing"
	"time"
)

func TestZSet_ZAdd(t *testing.T) {
	store, err := NewKVStore("pebble", "test")
	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	z := ZSet{kvstore: store, pk: "hello"}

	tbl := []struct {
		subKey string
		score  float64
		rank   int64
	}{
		{
			"subKey1",
			80,
			1,
		},
		{
			"subkey2",
			90,
			3,
		},

		{
			"subkey3",
			100,
			2,
		},
	}

	for _, v := range tbl {
		err := z.ZAdd(v.subKey, v.score)
		if err != nil {
			t.Fatal(err)
		}
	}

	err = z.ZAdd("subKey1", 200)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range tbl {
		rank := z.ZRank(v.subKey, true)
		if rank != v.rank {
			t.Errorf("subKey %s,rank:%d expected:%v", v.subKey, rank, v.rank)
		}
	}

	rank := z.ZRank("subKey4", true)
	if rank != -1 {
		t.Errorf("subKey %s,rank:%d expected:%v", "subKey4", rank, -1)
	}

}

func TestZSet_ZRange(t *testing.T) {
	store, err := NewKVStore("pebble", "test")
	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	z := ZSet{kvstore: store, pk: "hello"}

	tbl := []struct {
		subKey string
		score  float64
		rank   int64
	}{
		{
			"subKey1",
			120,
			1,
		},
		{
			"subkey2",
			90,
			3,
		},

		{
			"subkey3",
			100,
			2,
		},
	}

	expect := []string{"subKey1", "subkey3", "subkey2"}

	for _, v := range tbl {
		z.ZAdd(v.subKey, v.score)
	}

	ls := z.ZRange(0, -1, true)

	for i, v := range ls {
		if expect[i] != v.subKey {
			t.Errorf("subKey %s,expected:%v", v.subKey, expect[i])
		}
	}
}

func BenchmarkZSet_ZAdd(b *testing.B) {
	store, err := NewKVStore("pebble", "test")
	if err != nil {
		b.Fatal(err)
	}

	//defer store.close()

	z1 := ZSet{kvstore: store, pk: "hello"}
	//z2 := ZSet{kvstore: store, pk: "hello1"}

	for i := 0; i < b.N; i++ {
		z1.ZAdd(toString(int32(i)), float64(i))
		//go z2.ZAdd(toString(int32(i)), float64(i))
	}

	time.Sleep(time.Second)
	store.Close()

}

func TestZSet_ZAddManual(t *testing.T) {
	store, err := NewKVStore("pebble", "test")
	if err != nil {
		t.Fatal(err)
	}

	defer store.Close()

	z1 := ZSet{kvstore: store, pk: "hello"}

	concurrent := 1

	var wg sync.WaitGroup
	wg.Add(concurrent)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
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

func Benchmark_Pebble(b *testing.B) {
	s, err := newPebbleScore("test1")
	if err != nil {
		b.Fatal(err)
	}
	defer s.close()

	for i := 0; i < b.N; i++ {
		data := []byte(toString(int32(i)))
		m := map[string][]byte{
			string(data): data,
		}
		err := s.commit(m, nil, nil)
		//err := s.db.Set(data, data, &pebble.WriteOptions{})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestZSet_DeleteZSet(t *testing.T) {
	base, err := newPebbleScore("test")
	if err != nil {
		t.Fatal(err)
	}

	store := NewKVStoreWithBase(base)
	defer store.Close()

	printKeyCount(base.db)

	z := ZSet{kvstore: store, pk: "hello"}
	now := time.Now()
	z.DeleteZSet()

	fmt.Println("del elapse", time.Since(now))
	printKeyCount(base.db)
}

func printKeyCount(db *pebble.DB) {
	iter := db.NewIter(&pebble.IterOptions{
		LowerBound:  nil,
		UpperBound:  nil,
		TableFilter: nil,
	})

	type ZSetMeta struct {
		name       string
		count      uint32
		hCount     int32
		sCount     int32
		totalCount int32
	}

	m := map[string]*ZSetMeta{}
	var errCount int32

	for iter.First(); iter.Valid(); iter.Next() {
		key := iter.Key()
		mk, err := decodeRecordKey(string(key))
		if err != nil {
			errCount++
			continue
		}

		if _, ok := m[mk.getPrimaryKey()]; !ok {
			m[mk.getPrimaryKey()] = &ZSetMeta{
				name:   mk.getPrimaryKey(),
				count:  0,
				hCount: 0,
				sCount: 0,
			}
		}

		switch mk.RecordType {
		case RT_ZSET_META:
			v, _ := decodeRecordValue(iter.Value())
			v1, _ := decodeZSLMetaValue(v.value)
			m[mk.getPrimaryKey()].count = v1.count - 1
		case RT_ZSET_S_ELE:
			m[mk.getPrimaryKey()].sCount++
		case RT_ZSET_H_ELE:
			m[mk.getPrimaryKey()].hCount++
		default:

		}

		m[mk.getPrimaryKey()].totalCount++
	}

	fmt.Println("errCount", errCount)

	for _, v := range m {
		fmt.Printf("name:%s count:%d hcount:%d scount:%d total:%d \n", v.name, v.count, v.hCount, v.sCount, v.totalCount)
	}
}
