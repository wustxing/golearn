package skiplist

import (
	"errors"
)

type KVStoreBase interface {
	getKV(key []byte) ([]byte, error)
	commit(updates map[string][]byte, dels [][]byte, delRange [][2][]byte) error
	close() error
}

type KVStore struct {
	base KVStoreBase
}

func NewKVStore(storeType string, name string) (*KVStore, error) {
	switch storeType {
	case "pebble":
		base, err := newPebbleScore(name)
		if err != nil {
			return nil, err
		}
		return &KVStore{base: base}, nil
	case "memory":
		base := newMemoryScore(name)
		return &KVStore{base: base}, nil
	default:
		return nil, errors.New("not support type")
	}
}

func NewKVStoreWithBase(base KVStoreBase) *KVStore {
	return &KVStore{base: base}
}

func (p *KVStore) getKV(key RecordKey) (RecordValue, bool) {
	data, err := p.base.getKV(key.encode())
	if err != nil {
		if err == errNotFound {
			return RecordValue{}, false
		}

		return RecordValue{}, false
	}

	v, err := decodeRecordValue(data)
	if err != nil {
		return RecordValue{}, false
	}
	return v, true
}

func (p *KVStore) createCacheStore() *CacheStore {
	return &CacheStore{cache: map[RecordKey]RecordValue{}, parent: p}
}

func (p *KVStore) Close() error {
	return p.base.close()
}

type CacheStore struct {
	parent    *KVStore
	cache     map[RecordKey]RecordValue
	delRanges [][2]RecordKey
}

func (p *CacheStore) delKV(key RecordKey) {
	p.cache[key] = RecordValue{
		RecordType: RT_Invalid,
	}
}

func (p *CacheStore) delRange(start, end RecordKey) {
	p.delRanges = append(p.delRanges, [2]RecordKey{start, end})
}

func (p *CacheStore) setKV(key RecordKey, value RecordValue) {
	p.cache[key] = value
}

func (p *CacheStore) getKV(key RecordKey) (RecordValue, bool) {
	v, ok := p.cache[key]
	if ok {
		if v.RecordType == RT_Invalid {
			return RecordValue{}, false
		}
		return v, true
	}

	v, ok = p.parent.getKV(key)
	return v, ok
}

func (p *CacheStore) commit() error {
	updates := make(map[string][]byte)
	dels := make([][]byte, 0)
	for k, v := range p.cache {
		if v.RecordType == RT_Invalid {
			dels = append(dels, k.encode())
			continue
		}
		updates[string(k.encode())] = v.encode()
	}

	var delRanges [][2][]byte
	for _, v := range p.delRanges {
		r := [2][]byte{v[0].encode(), v[1].encode()}
		delRanges = append(delRanges, r)
	}

	return p.parent.base.commit(updates, dels, delRanges)
}
