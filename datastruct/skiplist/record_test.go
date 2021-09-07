package skiplist

import (
	"fmt"
	"testing"
)

func Test_Varint(t *testing.T) {
	var a uint64 = 999999
	data := varintEncode(a)
	b, bsize := varintDecodeFwd(data, 10)
	if a != b {
		t.FailNow()
	}

	var dataRvs []byte
	for i := len(data) - 1; i >= 0; i-- {
		dataRvs = append(dataRvs, data[i])
	}

	c, csize := varintDecodeRvs(dataRvs, 10)
	if c != a {
		t.FailNow()
	}

	if bsize != csize {
		t.FailNow()
	}
}

func TestRecordKey(t *testing.T) {
	key := RecordKey{
		RecordType: RT_ZSET_META,
		pk:         "hello",
		sk:         "",
		version:    3,
	}
	data := key.encode()
	fmt.Println(string(data), data)

	key1, err := decodeRecordKey(string(data))
	if err != nil {
		t.Fatal(err)
	}

	if key != key1 {
		t.Fatal("key not equal")
	}
}

func TestZSLMetaValue(t *testing.T) {
	value := ZSLMetaValue{
		level:    3,
		maxLevel: 4,
		count:    5,
		tail:     6,
		posAlloc: 7,
	}

	data := value.encode()

	value1, err := decodeZSLMetaValue(data)
	if err != nil {
		t.Fatal(err)
	}

	if value != value1 {
		t.Fatal()
	}
}
