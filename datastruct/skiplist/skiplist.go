package skiplist

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

const (
	MaxLevel = 32
	HEAD_ID  = 1
)

//var gCache map[uint64]*ZSLEleValue
//
//func init() {
//	gCache = make(map[uint64]*ZSLEleValue)
//}

type Skiplist struct {
	posAlloc uint64

	level           uint8
	getFromCacheCnt int32
	getFromStore    int32
	deletedCnt      int32
	updatedCnt      int32
	cache           map[uint64]*ZSLEleValue
	tail            uint64

	count uint32

	pk string

	store *CacheStore
}

func NewSkiplist(pk string, meta ZSLMetaValue, store *CacheStore) *Skiplist {
	return &Skiplist{
		posAlloc:        meta.getPosAlloc(),
		level:           meta.getLevel(),
		getFromCacheCnt: 0,
		deletedCnt:      0,
		updatedCnt:      0,
		cache:           map[uint64]*ZSLEleValue{},
		tail:            meta.getTail(),
		count:           meta.getCount(),
		pk:              pk,
		store:           store,
	}
}

func (p *Skiplist) randomLevel() uint8 {
	level := uint8(1)
	for float32(rand.Int31()&0xFFFF) < (0.25 * 0xFFFF) {
		level++
	}
	if level < MaxLevel {
		return level
	}
	return MaxLevel
}

func (p *Skiplist) makeNode(score float64, key string) (uint64, *ZSLEleValue) {
	p.posAlloc++
	e := NewZSLEleValue(score, key)
	return p.posAlloc, &e
}

func (p *Skiplist) getEleByRank(rank uint32) (*ZSLEleValue, bool) {
	_, ok := p.getNode(HEAD_ID)
	if !ok {
		return nil, false
	}

	pos := uint64(HEAD_ID)
	traversed := uint32(0)

	for i := p.level; i >= 1; i-- {
		tmpPos := p.cache[pos].getForward(i)
		for tmpPos != 0 {
			next, ok := p.getNode(tmpPos)
			if !ok {
				return nil, false
			}

			if traversed+p.cache[pos].getSpan(i) <= rank {
				traversed += p.cache[pos].getSpan(i)
				pos = tmpPos
				tmpPos = next.getForward(i)
			} else {
				break
			}
		}
		if traversed == rank {
			return p.cache[pos], true
		}
	}

	return nil, false
}

func (p *Skiplist) getNode(pointer uint64) (*ZSLEleValue, bool) {
	if v, ok := p.cache[pointer]; ok {
		p.getFromCacheCnt++
		return v, true
	}

	//if v, ok := gCache[pointer]; ok {
	//	p.getFromCacheCnt++
	//	p.cache[pointer] = v
	//	return v, true
	//}

	rk := RecordKey{
		RecordType: RT_ZSET_S_ELE,
		pk:         p.pk,
		sk:         toStringInt64(pointer),
		version:    0,
	}

	rv, ok := p.store.getKV(rk)
	if !ok {
		return nil, false
	}

	result, err := decodeZSLEleValue(rv.value)
	if err != nil {
		return nil, false
	}

	p.cache[pointer] = &result
	//gCache[pointer] = &result
	p.getFromStore++
	return &result, true
}

func (p *Skiplist) delNode(pos uint64) {
	delete(p.cache, pos)
	//delete(gCache, pos)

	p.deletedCnt++

	rk := RecordKey{
		RecordType: RT_ZSET_S_ELE,
		pk:         p.pk,
		sk:         toStringInt64(pos),
		version:    0,
	}
	p.store.delKV(rk)
}

func (p *Skiplist) saveNode(pos uint64, val *ZSLEleValue) {
	p.cache[pos].setChanged(false)
	p.updatedCnt++

	rk := RecordKey{
		RecordType: RT_ZSET_S_ELE,
		pk:         p.pk,
		sk:         toStringInt64(pos),
		version:    0,
	}
	rv := RecordValue{
		RecordType: RT_ZSET_S_ELE,
		value:      val.encode(),
	}
	p.store.setKV(rk, rv)
}

func (p *Skiplist) save() {
	for pos, v := range p.cache {
		if v.isChanged() {
			p.saveNode(pos, v)
		}
	}

	rk := RecordKey{
		RecordType: RT_ZSET_META,
		pk:         p.pk,
		sk:         "",
		version:    0,
	}

	mv := ZSLMetaValue{
		level:    p.level,
		maxLevel: MaxLevel,
		count:    p.count,
		tail:     p.tail,
		posAlloc: p.posAlloc,
	}

	rv := RecordValue{
		RecordType: RT_ZSET_META,
		value:      mv.encode(),
	}

	p.store.setKV(rk, rv)
}

func (p *Skiplist) remove(score float64, subKey string) error {
	update := make([]uint64, MaxLevel+1)

	_, ok := p.getNode(HEAD_ID)
	if !ok {
		return errors.New("getNode")
	}

	pos := uint64(HEAD_ID)

	for i := p.level; i >= 1; i-- {
		tmpPos := p.cache[pos].getForward(i)
		for tmpPos != 0 {
			next, ok := p.getNode(tmpPos)
			if !ok {
				return errors.New("getNode")
			}

			if slCmp(next.getScore(), next.getSubKey(), score, subKey) < 0 {
				pos = tmpPos
				tmpPos = next.getForward(i)
			} else {
				break
			}
		}
		update[i] = pos
	}

	delPos := p.cache[pos].getForward(1)

	if delPos == 0 || slCmp(p.cache[delPos].getScore(), p.cache[delPos].getSubKey(), score, subKey) != 0 {
		return errors.New("not equal")
	}

	return p.removeInterval(delPos, update)
}

func (p *Skiplist) removeInterval(pos uint64, update []uint64) error {
	for i := uint8(1); i <= p.level; i++ {
		toupdate := p.cache[update[i]]
		if toupdate.getForward(i) != pos {
			toupdate.setSpan(i, toupdate.getSpan(i)-1)
		} else {
			toupdate.setSpan(i, toupdate.getSpan(i)+p.cache[pos].getSpan(i)-1)
			toupdate.setForward(i, p.cache[pos].getForward(i))
		}
	}

	btmFwd := p.cache[pos].getForward(1)
	if btmFwd > 0 {
		node, ok := p.getNode(btmFwd)
		if !ok {
			return errors.New("getNode")
		}

		node.setBackward(p.cache[pos].getBackward())
	} else {
		p.tail = p.cache[pos].getBackward()
	}

	p.count--
	for p.level > 1 && p.cache[HEAD_ID].getForward(p.level) == 0 {
		p.level--
	}

	p.delNode(pos)
	return nil
}

func (p *Skiplist) removeRangeByRank(start, end uint32) bool {
	_, ok := p.getNode(HEAD_ID)
	if !ok {
		return false
	}

	update := make([]uint64, MaxLevel)
	pos := uint64(HEAD_ID)
	traversed := uint32(0)

	for i := p.level; i >= 1; i-- {
		tmpPos := p.cache[pos].getForward(i)
		for tmpPos != 0 {
			next, ok := p.getNode(tmpPos)
			if !ok {
				return false
			}

			if traversed+p.cache[pos].getSpan(i) < start {
				traversed += p.cache[pos].getSpan(i)
				pos = tmpPos
				tmpPos = next.getForward(i)
			} else {
				break
			}
		}
		update[i] = pos
	}

	traversed += 1

	pos = p.cache[pos].getForward(1)
	for pos > 0 && traversed < end {
		traversed += 1
		next := p.cache[pos].getForward(1)
		err := p.removeInterval(pos, update)
		if err != nil {
			return false
		}
		pos = next
		if pos != 0 {
			_, ok := p.getNode(pos)
			if !ok {
				return false
			}
		}
	}
	return true
}

func (p *Skiplist) rank(score float64, key string) (uint32, bool) {
	var rank uint32 = 0

	x, ok := p.getNode(HEAD_ID)
	if !ok {
		return 0, false
	}

	for i := p.level; i >= 1; i-- {
		tmpPos := x.getForward(i)
		for tmpPos > 0 {
			next, ok := p.getNode(tmpPos)
			if !ok {
				return 0, false
			}

			cmp := slCmp(next.getScore(), next.getSubKey(), score, key)
			if cmp <= 0 {
				rank += x.getSpan(i)
				x = next
				tmpPos = x.getForward(i)
			} else {
				break
			}
		}
	}

	if rank != 0 && x.getSubKey() == key {
		return rank, true
	}
	return 0, false
}

func slCmp(score0 float64,
	subk0 string,
	score1 float64,
	subk1 string) int {
	if (score0 == score1) && (subk0 == subk1) {
		return 0
	}
	if (score0 < score1) || (score0 == score1 && subk0 < subk1) {
		return -1
	}
	if (score0 > score1) || (score0 == score1 && subk0 > subk1) {
		return 1
	}
	return 0
}

func (p *Skiplist) insert(score float64, subkey string) error {
	if p.count >= math.MaxInt32/2 {
		return errors.New("zset count reach limit")
	}

	update := make([]uint64, MaxLevel+1)
	rank := make([]uint32, MaxLevel+1)

	_, ok := p.getNode(HEAD_ID)
	if !ok {
		return errors.New("get node")
	}

	pos := uint64(HEAD_ID)

	for i := p.level; i >= 1; i-- {
		tmpPos := p.cache[pos].getForward(i)
		if i != p.level {
			// accumulate upper level's rank
			rank[i] = rank[i+1]
		}
		for tmpPos != 0 {
			// TODO(deyukong): get from cache first
			next, ok := p.getNode(tmpPos)
			if !ok {
				return errors.New("get node")
			}

			if next.getSubKey() == subkey {
				return errors.New("duplicate key")
			}

			if slCmp(next.getScore(), next.getSubKey(), score, subkey) < 0 {
				rank[i] += p.cache[pos].getSpan(i)
				pos = tmpPos
				tmpPos = next.getForward(i)
			} else {
				break
			}
		}
		update[i] = pos
	}

	lvl := p.randomLevel()
	if lvl > p.level {
		for i := p.level + 1; i <= lvl; i++ {
			rank[i] = 0
			update[i] = HEAD_ID
			// NOTE(deyukong): head node also affects _count, so here the span
			// should be _count -1, not _count.
			p.cache[update[i]].setSpan(i, p.count-1)
		}
		p.level = lvl
	}

	pos, node := p.makeNode(score, subkey)
	p.cache[pos] = node
	//gCache[pos] = node

	for i := uint8(1); i <= lvl; i++ {
		//INVARIANT(update[i] >= ZSlMetaValue::HEAD_ID);
		//INVARIANT(cache.find(update[i]) != cache.end());
		p.cache[pos].setForward(i, p.cache[update[i]].getForward(i))
		p.cache[update[i]].setForward(i, pos)
		p.cache[pos].setSpan(i,
			p.cache[update[i]].getSpan(i)-(rank[1]-rank[i]))
		p.cache[update[i]].setSpan(i, rank[1]-rank[i]+1)
	}
	for i := lvl + 1; i <= p.level; i++ {
		p.cache[update[i]].setSpan(i, p.cache[update[i]].getSpan(i)+1)
	}

	if update[1] == HEAD_ID {
		p.cache[pos].setBackward(0)
	} else {
		p.cache[pos].setBackward(update[1])
	}

	btmFwd := p.cache[pos].getForward(1)
	if btmFwd != 0 {
		_, ok := p.getNode(btmFwd)
		if !ok {
			return errors.New("get node")
		}
		p.cache[btmFwd].setBackward(pos)
	} else {
		p.tail = pos
	}
	p.count++
	return nil
}

func (p *Skiplist) getCount() uint32 {
	return p.count
}

func (p *Skiplist) scanByRank(start int64, len int32, rev bool) []*ZSLEleValue {
	var ln *ZSLEleValue
	if rev {
		expTail, ok := p.getNode(p.tail)
		if !ok {
			return nil
		}

		ln = expTail
		if start > 0 {
			tmp, ok := p.getEleByRank(p.count - 1 - uint32(start))
			if !ok {
				return nil
			}
			ln = tmp
		}
	} else {
		expHead, ok := p.getNode(HEAD_ID)
		if !ok {
			return nil
		}
		first := expHead.getForward(1)
		expNode, ok := p.getNode(first)
		if !ok {
			return nil
		}
		ln = expNode
		if start > 0 {
			tmp, ok := p.getEleByRank(uint32(start) + 1)
			if !ok {
				return nil
			}
			ln = tmp
		}
	}
	if ln == nil {
		return nil
	}
	result := make([]*ZSLEleValue, 0)

	for {
		len--
		//INVARIANT(ln != nullptr);
		// std::cout << ln->getScore() << ' ' << ln->getSubKey() << std::endl;
		result = append(result, ln)
		if len == 0 {
			break
		}

		if rev {
			tmp, ok := p.getNode(ln.getBackward())

			if !ok {
				return nil
			}
			ln = tmp
		} else {
			tmp, ok := p.getNode(ln.getForward(1))
			if !ok {
				return nil
			}
			ln = tmp
		}
	}
	return result
}

func (p *Skiplist) PrintStatus() {
	fmt.Printf("update:%d del:%d getFromStore:%d getFromCache:%d \n", p.updatedCnt, p.deletedCnt, p.getFromStore, p.getFromCacheCnt)
}
