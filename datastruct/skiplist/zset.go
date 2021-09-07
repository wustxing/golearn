package skiplist

import (
	"errors"
	"math"
	"sync"
)

type ZSet struct {
	kvstore *KVStore
	pk      string
	sync.RWMutex
}

func NewZSet(store *KVStore, pk string) *ZSet {
	return &ZSet{
		kvstore: store,
		pk:      pk,
		RWMutex: sync.RWMutex{},
	}
}

func (p *ZSet) ZAdd(subKey string, score float64) error {
	p.Lock()
	defer p.Unlock()

	pk := p.pk
	store := p.kvstore.createCacheStore()

	mk, meta, err := getZSetPKKeyMeta(pk, p.kvstore)

	if err != nil {
		if err != errNotFound {
			return err
		}
	}

	if err == errNotFound {
		mk = RecordKey{
			RecordType: RT_ZSET_META,
			pk:         pk,
			sk:         "",
			version:    0,
		}

		tmp := ZSLMetaValue{
			level:    1,
			maxLevel: MaxLevel,
			count:    1,
			tail:     0,
			posAlloc: 128,
		}
		rv := RecordValue{
			RecordType: RT_ZSET_META,
			value:      tmp.encode(),
		}

		store.setKV(mk, rv)

		head := RecordKey{
			RecordType: RT_ZSET_S_ELE,
			pk:         mk.getPrimaryKey(),
			sk:         toString(HEAD_ID),
			version:    0,
		}

		headVal := NewZSLEleValue(0, "")

		subRv := RecordValue{
			RecordType: RT_ZSET_S_ELE,
			value:      headVal.encode(),
		}

		store.setKV(head, subRv)
		meta = tmp
	}

	sl := NewSkiplist(mk.getPrimaryKey(), meta, store)

	hk := RecordKey{
		RecordType: RT_ZSET_H_ELE,
		pk:         mk.getPrimaryKey(),
		sk:         subKey,
		version:    0,
	}

	eValue, ok := store.getKV(hk)
	if !ok {
		sl.insert(score, subKey)
		hv := newRecordValueFloat64(score, RT_ZSET_H_ELE)
		store.setKV(hk, hv)
	} else {
		oldScore, ok := float64Decode(eValue.value)
		if !ok {
			return errors.New("float64Decode")
		}
		err := sl.remove(oldScore, subKey)
		if err != nil {
			return err
		}

		err = sl.insert(score, subKey)
		if err != nil {
			return err
		}

		hv := newRecordValueFloat64(score, RT_ZSET_H_ELE)
		store.setKV(hk, hv)
	}

	sl.save()
	//sl.PrintStatus()
	//return nil
	return store.commit()
}

func (p *ZSet) ZRank(subKey string, reverse bool) int64 {
	p.RLock()
	defer p.RUnlock()

	pk := p.pk

	store := p.kvstore.createCacheStore()

	_, meta, err := getZSetPKKeyMeta(pk, p.kvstore)
	if err != nil {
		return -1
	}

	hk := RecordKey{
		RecordType: RT_ZSET_H_ELE,
		pk:         pk,
		sk:         subKey,
		version:    0,
	}

	eValue, ok := store.getKV(hk)
	if !ok {
		return -1
	}

	score, ok := float64Decode(eValue.value)
	if !ok {
		return -1
	}

	sl := NewSkiplist(pk, meta, store)

	rank, ok := sl.rank(score, subKey)
	if !ok {
		return -1
	}

	r := rank
	if reverse {
		r = meta.getCount() - r
	}
	return int64(r)
}

func (p *ZSet) ZRange(start, end int64, rev bool) []*ZSLEleValue {
	p.RLock()
	defer p.RUnlock()

	pk := p.pk

	store := p.kvstore.createCacheStore()

	_, meta, err := getZSetPKKeyMeta(pk, p.kvstore)
	if err != nil {
		return nil
	}
	sl := NewSkiplist(pk, meta, store)
	length := int64(sl.getCount() - 1)

	if start < 0 {
		start = length + start
	}

	if end < 0 {
		end = length + end
	}

	if start > end || start >= length {
		return nil
	}

	if end >= length {
		end = length - 1
	}

	rangeLen := end - start + 1

	arr := sl.scanByRank(start, int32(rangeLen), rev)
	return arr
}

func (p *ZSet) DeleteZSet() error {
	p.Lock()
	defer p.Unlock()

	pk := p.pk

	mk, _, err := getZSetPKKeyMeta(pk, p.kvstore)
	if err != nil {
		return err
	}

	cacheStore := p.kvstore.createCacheStore()
	cacheStore.delKV(mk)
	delSubKeysRange(mk, cacheStore)
	return cacheStore.commit()
}

var errNotFound = errors.New("not found")
var errDecode = errors.New("decord error")

func getZSetPKKeyMeta(pk string, store *KVStore) (RecordKey, ZSLMetaValue, error) {
	key := RecordKey{
		RecordType: RT_ZSET_META,
		pk:         pk,
		sk:         "",
		version:    0,
	}

	eValue, ok := store.getKV(key)
	if !ok {
		return RecordKey{}, ZSLMetaValue{}, errNotFound
	}

	meta, err := decodeZSLMetaValue(eValue.value)
	if err != nil {
		return RecordKey{}, ZSLMetaValue{}, errDecode
	}

	return key, meta, nil
}

func delSubKeysRange(mk RecordKey, store *CacheStore) {
	start := RecordKey{
		RecordType: RT_ZSET_S_ELE,
		pk:         mk.getPrimaryKey(),
		sk:         "",
		version:    0,
	}

	end := RecordKey{
		RecordType: RT_ZSET_S_ELE,
		pk:         mk.getPrimaryKey(),
		sk:         "",
		version:    math.MaxUint64,
	}

	store.delRange(start, end)

	start = RecordKey{
		RecordType: RT_ZSET_H_ELE,
		pk:         mk.getPrimaryKey(),
		sk:         "",
		version:    0,
	}

	end = RecordKey{
		RecordType: RT_ZSET_H_ELE,
		pk:         mk.getPrimaryKey(),
		sk:         "",
		version:    math.MaxUint64,
	}

	store.delRange(start, end)
	return
}
