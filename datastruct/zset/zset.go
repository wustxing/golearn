package zset

/* Input flags. */
const ZADD_NONE = 0
const ZADD_INCR = (1 << 0) /* Increment the Score instead of setting it. */
const ZADD_NX = (1 << 1)   /* Don't touch elements not already existing. */
const ZADD_XX = (1 << 2)   /* Only touch elements already existing. */

/* Output flags. */
const ZADD_NOP = (1 << 3)     /* Operation not performed because of conditionals.*/
const ZADD_NAN = (1 << 4)     /* Only touch elements already existing. */
const ZADD_ADDED = (1 << 5)   /* The element was new and was added. */
const ZADD_UPDATED = (1 << 6) /* The element already existed, Score updated. */

/* Flags only used by the ZADD command but not by zsetAdd() API: */
const ZADD_CH = (1 << 16) /* Return num of elements added or updated. */

type ZSet struct {
	zsl  *skipList
	dict map[string]DictV
}

type DictV struct {
	Score int64
	Extra string
}

type Element struct {
	Member string
	Score  int64
	Extra  string //额外参数
}

func New() *ZSet {
	p := &ZSet{}
	p.zsl = newSkipList()
	p.dict = make(map[string]DictV, 0)
	return p
}

func (p *ZSet) Length() int {
	return p.zsl.length
}

func (p *ZSet) Score(ele string) (int64, bool) {
	v, exist := p.dict[ele]
	if exist {
		return v.Score, true
	}
	return 0, false
}

func (p *ZSet) V(ele string) (DictV, bool) {
	v, exist := p.dict[ele]
	return v, exist
}

func (p *ZSet) add(ele string, score int64, extra string, flags int) (flagsRet int, newScore int64) {
	incr := (flags & ZADD_INCR) != 0
	nx := (flags & ZADD_NX) != 0
	xx := (flags & ZADD_XX) != 0
	var curScore int64

	de, exist := p.dict[ele]
	if exist {
		if nx {
			flagsRet |= ZADD_NOP
			return
		}
		curScore = de.Score
		if incr {
			score += curScore
			newScore = score
		}

		if score != curScore {
			znode := p.zsl.update(ele, curScore, score, extra)
			p.dict[ele] = DictV{
				Score: znode.score,
				Extra: znode.extra,
			}
			flagsRet |= ZADD_UPDATED
			newScore = score
		}
		return
	} else if !xx {
		znode := p.zsl.insert(ele, score, extra)
		p.dict[ele] = DictV{
			Score: znode.score,
			Extra: znode.extra,
		}
		flagsRet |= ZADD_ADDED
		newScore = score
		return
	} else {
		flagsRet |= ZADD_NOP
		return
	}
	return
}

func (p *ZSet) Rank(ele string, reverse bool) int {
	v, exist := p.dict[ele]
	if !exist {
		return -1
	}

	rank := p.zsl.GetRank(ele, v.Score)
	if reverse {
		return p.Length() - int(rank)
	} else {
		return int(rank - 1)
	}
}

func (p *ZSet) Range(start int, end int, reverse bool) (rets []*Element) {
	llen := p.Length()

	if start < 0 {
		start = llen + start
	}
	if end < 0 {
		end = llen + end
	}
	if start < 0 {
		start = 0
	}

	if start > end || start >= llen {
		return
	}

	if end >= llen {
		end = llen - 1
	}

	//	rangeLen := (end - start) + 1
	var ln *skiplistNode
	if reverse {
		ln = p.zsl.tail
		if start > 0 {
			ln = p.zsl.GetElementByRank(uint(llen - start))
		}
	} else {
		ln = p.zsl.header.level[0].forward
		if start > 0 {
			ln = p.zsl.GetElementByRank(uint(start + 1))
		}
	}

	for i := (end - start) + 1; i > 0; i-- {
		rets = append(rets, &Element{
			Member: ln.member,
			Score:  ln.score,
			Extra:  "",
		})
		if reverse {
			ln = ln.backward
		} else {
			ln = ln.level[0].forward
		}
	}
	return
}

//返回更新了几列
func (p *ZSet) AddBatch(datas ...Element) (changed int) {
	changed, _ = p.addGeneric(ZADD_NONE, datas...)
	return
}

func (p *ZSet) Add(ele string, score int64, params ...string) int64 {
	var extra string
	if len(params) > 0 {
		extra = params[0]
	}
	d := Element{
		Member: ele,
		Score:  score,
		Extra:  extra,
	}
	_, newScore := p.addGeneric(ZADD_NONE, d)
	return newScore
}

//返回更新后的分
func (p *ZSet) Incr(ele string, incr int64) (score int64) {
	_, score = p.addGeneric(ZADD_INCR, Element{
		Member: ele,
		Score:  incr,
		Extra:  "",
	})
	return
}

//当incr时，newScore才是有效的
func (p *ZSet) addGeneric(flags int, datas ...Element) (changed int, newScore int64) {
	incr := (flags & ZADD_INCR) != 0
	nx := (flags & ZADD_NX) != 0
	xx := (flags & ZADD_XX) != 0
	//ch := (flags & ZADD_CH) != 0

	if nx && xx {
		return
	}

	if incr && len(datas) > 1 {
		return
	}

	var added, updated, processed int
	var score int64
	for _, data := range datas {
		retflags, newScore := p.add(data.Member, data.Score, data.Extra, flags)
		if (retflags & ZADD_ADDED) != 0 {
			added++
		}
		if (retflags & ZADD_UPDATED) != 0 {
			updated++
		}
		if (retflags & ZADD_NOP) == 0 {
			processed++
		}
		score = newScore
	}

	return (added + updated), score
}
