package skiplist

import (
	"errors"
	"fmt"
)

type RecordType int32

const (
	RT_Invalid RecordType = iota
	RT_ZSET_META
	RT_ZSET_S_ELE
	RT_ZSET_H_ELE
)

type RecordKey struct {
	RecordType
	pk      string
	sk      string
	version uint64
}

func (p *RecordKey) getPrimaryKey() string {
	return p.pk
}

func (p *RecordKey) prefixPk() (ret []byte) {
	ret = make([]byte, 0, 128)
	ret = append(ret, rt2Char(p.RecordType))
	ret = append(ret, []byte(p.pk)...)
	ret = append(ret, 0)
	ret = append(ret, varintEncode(p.version)...)
	return ret
}

func (p *RecordKey) encode() (ret []byte) {
	ret = p.prefixPk()
	ret = append(ret, p.sk...)

	lenPK := varintEncode(uint64(len(p.pk)))

	for i := len(lenPK) - 1; i >= 0; i-- {
		ret = append(ret, lenPK[i])
	}
	return
}

func decodeRecordKey(key string) (RecordKey, error) {
	offset := 1

	recordType := decodeTye(key)

	pkLen64, pkLenSize := varintDecodeRvs([]byte(key), uint64(len(key)))

	pkLen := int(pkLen64)

	if len(key) < pkLen+offset {
		return RecordKey{}, errors.New("invalid pklen")
	}

	pk := key[offset : pkLen+offset]

	ptr := offset + pkLen + 1
	version, verSize := varintDecodeFwd([]byte(key)[ptr:], uint64(len(key)))

	left := len(key) - offset - pkLen - 1 - verSize - pkLenSize

	var sk string
	if left > 0 {
		sk = key[ptr+verSize : ptr+verSize+left]
	}
	return RecordKey{
		RecordType: recordType,
		pk:         pk,
		sk:         sk,
		version:    version,
	}, nil
}

func (p *RecordKey) getHdrSize() int32 {
	return 1
}

func decodeTye(key string) RecordType {
	if len(key) == 0 {
		return RT_Invalid
	}
	return char2Rt(key[0])
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

func varintDecodeFwd(input []byte, maxSize uint64) (uint64, int) {
	var ret uint64 = 0
	var i uint64 = 0
	for ; i < maxSize && (input[i]&0x80) > 0; i++ {
		if len(input) <= int(i) {
			fmt.Println("error")
		}
		ret |= uint64(input[i]&0x7f) << (7 * i)
	}
	if i == maxSize {
		return 0, 0
	}
	ret |= uint64(input[i]&0x7f) << (7 * (i))
	i++
	return ret, int(i)
}

func varintDecodeRvs(input []byte, maxSize uint64) (uint64, int) {

	var ret uint64 = 0

	var i uint64 = 0
	for j := uint64(len(input) - 1); j >= 0; j-- {

		b := input[j]

		ret |= uint64(b&0x7f) << (7 * i)

		i++
		if i >= maxSize || b&0x80 == 0 {
			break
		}
	}

	if i == maxSize {
		return 0, 0
	}

	//ret |= uint64(input[i]&0x7f) << (7 * (i))
	//i++
	return ret, int(i)
}

func rt2Char(t RecordType) byte {
	switch t {
	case RT_ZSET_META:
		return 'Z'
	case RT_ZSET_S_ELE:
		return 'c'
	case RT_ZSET_H_ELE:
		return 'z'
	default:
		return 0
	}
}

func char2Rt(t byte) RecordType {
	switch t {
	case 'Z':
		return RT_ZSET_META
	case 'c':
		return RT_ZSET_S_ELE
	case 'z':
		return RT_ZSET_H_ELE
	default:
		return RT_Invalid
	}
}

type RecordValue struct {
	RecordType
	value []byte
}

func (p *RecordValue) encode() []byte {
	ret := make([]byte, 0)
	ret = append(ret, rt2Char(p.RecordType))
	ret = append(ret, p.value...)
	return ret
}

func newRecordValueFloat64(v float64, recordType RecordType) RecordValue {
	d := float64Encode(v)
	return RecordValue{
		RecordType: recordType,
		value:      d,
	}
}

func decodeRecordValue(value []byte) (RecordValue, error) {
	if len(value) == 0 {
		return RecordValue{}, errors.New("value len==0")
	}

	recordType := char2Rt(value[0])

	offset := 1

	return RecordValue{
		RecordType: recordType,
		value:      value[offset:],
	}, nil
}

type ZSLMetaValue struct {
	level    uint8
	maxLevel uint8
	count    uint32
	tail     uint64
	posAlloc uint64
}

func (p *ZSLMetaValue) getMaxLevel() uint8 {
	return p.maxLevel
}

func (p *ZSLMetaValue) getLevel() uint8 {
	return p.level
}

func (p *ZSLMetaValue) getCount() uint32 {
	return p.count
}

func (p *ZSLMetaValue) getTail() uint64 {
	return p.tail
}

func (p *ZSLMetaValue) getPosAlloc() uint64 {
	return p.posAlloc
}

func (p *ZSLMetaValue) encode() []byte {
	value := make([]byte, 0, 128)

	value = append(value, varintEncode(uint64(p.level))...)
	value = append(value, varintEncode(uint64(p.maxLevel))...)
	value = append(value, varintEncode(uint64(p.count))...)
	value = append(value, varintEncode(uint64(p.tail))...)
	value = append(value, varintEncode(uint64(p.posAlloc))...)

	return value
}

func decodeZSLMetaValue(val []byte) (ZSLMetaValue, error) {
	var offset int = 0

	level, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	offset += size

	maxLevel, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	offset += size

	count, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	offset += size

	tail, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	offset += size

	posAlloc, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	offset += size

	return ZSLMetaValue{
		level:    uint8(level),
		maxLevel: uint8(maxLevel),
		count:    uint32(count),
		tail:     tail,
		posAlloc: posAlloc,
	}, nil
}

type ZSLEleValue struct {
	forward []uint64
	span    []uint32

	score float64

	backward uint64

	changed bool

	subKey string
}

func NewZSLEleValue(score float64, subKey string) ZSLEleValue {
	return ZSLEleValue{
		forward:  make([]uint64, MaxLevel+1),
		span:     make([]uint32, MaxLevel+1),
		score:    score,
		backward: 0,
		changed:  false,
		subKey:   subKey,
	}
}

func (p *ZSLEleValue) getBackward() uint64 {
	return p.backward
}

func (p *ZSLEleValue) setBackward(pointer uint64) {
	p.changed = true
	p.backward = pointer
}

func (p *ZSLEleValue) getForward(layer uint8) uint64 {
	return p.forward[layer]
}

func (p *ZSLEleValue) setForward(layer uint8, pointer uint64) {
	p.changed = true
	p.forward[layer] = pointer
}

func (p *ZSLEleValue) getSpan(layer uint8) uint32 {
	if p == nil {
		fmt.Println(layer)
	}
	if layer >= uint8(len(p.span)) || layer < 0 {
		fmt.Println(layer)
	}
	return p.span[layer]
}

func (p *ZSLEleValue) setSpan(layer uint8, span uint32) {
	p.changed = true
	p.span[layer] = span
}

func (p *ZSLEleValue) getScore() float64 {
	return p.score
}

func (p *ZSLEleValue) getSubKey() string {
	return p.subKey
}

func (p *ZSLEleValue) isChanged() bool {
	return p.changed
}
func (p *ZSLEleValue) setChanged(v bool) {
	p.changed = v
}

func (p *ZSLEleValue) encode() []byte {
	values := make([]byte, 0, 128)

	for _, v := range p.forward {
		values = append(values, varintEncode(v)...)
	}

	for _, v := range p.span {
		values = append(values, varintEncode(uint64(v))...)
	}

	values = append(values, float64Encode(p.score)...)

	values = append(values, varintEncode(p.backward)...)

	values = append(values, varintEncode(uint64(len(p.subKey)))...)

	values = append(values, p.subKey...)
	return values
}

func decodeZSLEleValue(val []byte) (ZSLEleValue, error) {
	offset := 0

	result := NewZSLEleValue(0, "")

	for i := 0; i <= MaxLevel; i++ {
		v, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
		result.forward[i] = v
		offset += size
	}

	for i := 0; i <= MaxLevel; i++ {
		v, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
		result.span[i] = uint32(v)
		offset += size
	}

	d, ok := float64Decode(val[offset:])
	if !ok {
		return ZSLEleValue{}, errors.New("float64Decode")
	}
	result.score = d
	offset += 8

	v, size := varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	result.backward = v
	offset += size

	v, size = varintDecodeFwd(val[offset:], uint64(len(val)-offset))
	keyLen := v
	offset += size

	result.subKey = string(val[offset : offset+int(keyLen)])
	return result, nil
}
