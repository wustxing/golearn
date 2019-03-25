package util

import (
	"math/rand"
	"sanguosha.com/games/sgs/framework/util"
	"sync"
	"time"
)

//var myRand = rand.New(rand.NewSource(time.Now().UnixNano()))

//var globalRand = New(&LockedSource{src: NewSource(1).(Source64)})
var myRand = rand.New(&LockdSource{src: rand.NewSource(time.Now().UnixNano()).(rand.Source64)})

type LockdSource struct {
	lk  sync.Mutex
	src rand.Source64
}

func (r *LockdSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

func (r *LockdSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

// RandNum ...
func RandNum(num int32) int32 {
	return myRand.Int31n(num)
}

func Rand31() int32 {
	return myRand.Int31()
}
func Rand63() int64 {
	return myRand.Int63()
}

//生成不重复的随机数，随机数范围0~max,数量num
func GenDiffRandomNum(num, max int) (outIntArr []int) {
	if num >= max {
		for i := 0; i < int(max); i++ {
			outIntArr = append(outIntArr, i)
		}
		return
	}
	var tmpArr []int
	for i := 0; i < int(max); i++ {
		tmpArr = append(tmpArr, 1)
	}

	var rnd int

	for i := 0; i < int(num); i++ {
		for {
			rnd = int(RandNum(int32(max)))
			if tmpArr[rnd] != -1 {
				break
			}
		}
		outIntArr = append(outIntArr, rnd)
		tmpArr[rnd] = -1
	}
	return
}

func RandomIndex(len int) []int {
	pollorder := make([]int, len)
	for i := 0; i < len; i++ {
		j := util.RandIntn(i + 1)
		pollorder[i] = pollorder[j]
		pollorder[j] = i
	}
	return pollorder
}
