package zset

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const MaxLevel = 32
const Probability = 0.25

func randomLevel() (level int) {
	for level = 1; rand.Float32() < Probability && level < MaxLevel; {
		level++
	}
	return
}

type nodeLevel struct {
	forward *skiplistNode
	span    uint
}

type skiplistNode struct {
	level []*nodeLevel

	backward *skiplistNode
	score    int64
	member   string
	extra    string
}

type skipList struct {
	header *skiplistNode
	tail   *skiplistNode
	length int
	level  int
}

func newNode(level int, ele string, score int64, extra string) *skiplistNode {
	levels := make([]*nodeLevel, level)
	for i := 0; i < level; i++ {
		levels[i] = &nodeLevel{}
	}
	return &skiplistNode{score: score, level: levels, member: ele, extra: extra}
}

func newSkipList() *skipList {
	rand.Seed(time.Now().UnixNano())
	p := &skipList{}
	p.level = 1
	p.length = 0
	p.header = newNode(MaxLevel, "", 0, "")
	for i := 0; i < MaxLevel; i++ {
		p.header.level[i].forward = nil
		p.header.level[i].span = 0
	}
	p.header.backward = nil
	p.tail = nil
	return p
}

func (p *skipList) insert(ele string, score int64, extra string) *skiplistNode {
	x := p.header

	update := make([]*skiplistNode, MaxLevel)
	rank := make([]uint, MaxLevel)
	//rank[i]即update[i]的排名
	//update[i]即截取时，要更新的前置结点
	for i := p.level - 1; i >= 0; i-- {
		if i == p.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		f := x.level[i].forward
		for f != nil && (f.score < score || (f.score == score && strings.Compare(f.member, ele) < 0)) {
			rank[i] += x.level[i].span
			x = f
			f = f.level[i].forward
		}
		update[i] = x
	}

	//不存在，新插入
	level := randomLevel()

	if level > p.level {
		for i := p.level; i < level; i++ {
			rank[i] = 0
			update[i] = p.header
			update[i].level[i].span = uint(p.length)
		}
		p.level = level
	}

	x = newNode(level, ele, score, extra)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < p.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == p.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		p.tail = x
	}
	p.length++
	return x
}

func (p *skipList) deleteNode(x *skiplistNode, update []*skiplistNode) {
	for i := 0; i < p.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span--
		}
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		p.tail = x.backward
	}

	for p.level > 1 && p.header.level[p.level-1].forward == nil {
		p.level--
	}
	p.length--
}

func (p *skipList) delete(ele string, score int64) int {
	x := p.header
	update := make([]*skiplistNode, MaxLevel)
	for i := p.level - 1; i >= 0; i-- {
		f := x.level[i].forward
		for f != nil && (f.score < score || (f.score == score && strings.Compare(f.member, ele) < 0)) {
			x = f
			f = f.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward

	if x != nil && x.score == score && strings.Compare(x.member, ele) == 0 {
		p.deleteNode(x, update)
		return 1
	}
	return 0
}

func (p *skipList) update(ele string, curScore int64, newScore int64, extra string) *skiplistNode {
	x := p.header
	update := make([]*skiplistNode, MaxLevel)
	for i := p.level - 1; i >= 0; i-- {
		f := x.level[i].forward
		for f != nil && (f.score < curScore || (f.score == curScore && strings.Compare(f.member, ele) < 0)) {
			x = f
			f = f.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward

	if !(x != nil && x.score == curScore && strings.Compare(x.member, ele) == 0) {
		return nil
	}

	//如果不需要移动
	if (x.backward == nil || x.backward.score < newScore) && (x.level[0].forward == nil || x.level[0].forward.score > newScore) {
		x.score = newScore
		return x
	}

	//如果需要移动，直接删除重加
	p.deleteNode(x, update)
	newNode := p.insert(ele, newScore, extra)
	return newNode
}

func (p *skipList) GetRank(ele string, score int64) uint {
	x := p.header
	var rank uint
	for i := p.level - 1; i >= 0; i-- {
		f := x.level[i].forward
		for f != nil && (f.score < score || (f.score == score && strings.Compare(f.member, ele) <= 0)) {
			rank += x.level[i].span
			x = f
			f = f.level[i].forward
		}
		if x.member != "" && strings.Compare(x.member, ele) == 0 {
			return rank
		}
	}
	return 0
}

func (p *skipList) Print() {
	fmt.Println()
	for i := p.level - 1; i >= 0; i-- {
		fmt.Printf("***level %d:", i+1)
		x := p.header
		for x.level[i].forward != nil {
			next := x.level[0].forward

			for next != nil {
				if next == x.level[i].forward {
					break
				}
				fmt.Print("   ")
				next = next.level[0].forward
			}
			fmt.Printf("%d %d ", x.level[i].forward.score, x.level[i].span)
			x = x.level[i].forward
		}
		fmt.Println()
	}
}

func (p *skipList) GetElementByRank(rank uint) *skiplistNode {
	x := p.header
	var traversed uint
	for i := p.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}
