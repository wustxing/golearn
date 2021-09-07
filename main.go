package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(4 << 20)
}

func randomLevel() uint8 {
	level := uint8(1)
	for float32(rand.Int31()&0xFFFF) < (0.25 * 0xFFFF) {
		level++
	}
	if level < 32 {
		return level
	}
	return level
}

func varintEncode(val uint64) (ret []byte) {

	for val >= 0x80 {
		a := val & 0x7f
		v := 0x80 | a
		ret = append(ret, byte(v))
		val >>= 7
	}

	ret = append(ret, byte(val))
	return
}

func varintDecodeFwd(input []byte, maxSize uint64) uint64 {
	var ret uint64 = 0
	var i uint64 = 0
	for ; i < maxSize && (input[i]&0x80) > 0; i++ {
		ret |= uint64(input[i]&0x7f) << (7 * i)
	}
	if i == maxSize {
		return 0
	}
	ret |= uint64(input[i]&0x7f) << (7 * (i))
	i++
	return ret
}

func varintDecodeRvs(input []byte, maxSize uint64) uint64 {

	var ret uint64 = 0

	var i uint64 = 0
	for i := uint64(len(input) - 1); i >= 0; i-- {
		b := input[i]
		if i > maxSize || b&0x80 == 0 {
			break
		}

		ret |= uint64(b&0x7f) << (7 * i)
	}

	if i == maxSize {
		return 0
	}
	ret |= uint64(input[i]&0x7f) << (7 * (i))
	i++
	return ret
}
