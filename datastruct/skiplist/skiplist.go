package skiplist

import (
	"fmt"
	"math/rand"
	"time"
)

const MaxLevel = 32
const Probability = 0.5

func randomLevel() (level int) {
	//rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; {
		level++
	}
	return
}

type nodeLevel struct {
	forward *node
	span    uint32
}

type node struct {
	level []*nodeLevel

	backward *node
	score    int
	obj      interface{}
}

type skipList struct {
	header *node
	tail   *node
	length uint32
	level  int
}

func newNode(score, level int) *node {
	return &node{score: score, level: make([]*nodeLevel, level)}
}

func newSkipList() *skipList {
	rand.Seed(time.Now().UnixNano())
	return &skipList{
		header: newNode(0, MaxLevel),
		level:  1,
	}
}

func (p *skipList) insert(score int, value interface{}) {
	x := p.header

	update := make([]*node, MaxLevel)
	rank := make([]uint32, MaxLevel)
	for i := p.level - 1; i >= 0; i-- {
		if i == p.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && x.level[i].forward.score < score {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	//如果存在此值，就替换
	x = x.level[0].forward
	if x != nil && x.score == score {
		x.obj = value
		return
	}

	//不存在，新插入
	level := randomLevel()
	if level > p.level {
		for i := p.level; i < level; i++ {
			rank[i] = 0
			update[i] = p.header
			update[i].level[i].span = p.length
		}
		p.level = level
	}

	node := newNode(score, level)
	for i := 0; i < level; i++ {
		node.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = node
		node.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
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
	if node.level[0].forward != nil {
		node.level[0].forward = node
	} else {
		p.tail = node
	}
	p.length++
	return
}

func (p *skipList) delete(key int) {
	//update := make([]*node, MaxLevel)
	//x := p.header
	//for i := p.level - 1; i >= 0; i-- {
	//	for x.forward[i] != nil && x.forward[i].key < key {
	//		x = x.forward[i]
	//	}
	//	update[i] = x
	//}
	//x = x.forward[0]
	//if x != nil && x.score == key {
	//	for i := 0; i < p.level; i++ {
	//		if update[i].forward[i] != x {
	//			break
	//		}
	//		update[i].forward[i] = x.forward[i]
	//	}
	//	for p.level > 1 && p.header.forward[p.level-1] == nil {
	//		p.level = p.level - 1
	//	}
	//}
}

func (p *skipList) search(key int) *node {
	//x := p.header
	//for i := p.level - 1; i >= 0; i-- {
	//	for x.forward[i] != nil && x.forward[i].key < key {
	//		x = x.forward[i]
	//		fmt.Printf("changeto %d\n", x.score)
	//	}
	//}
	//x = x.forward[0]
	//if x != nil && x.score == key {
	//	return x
	//}
	return nil
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
