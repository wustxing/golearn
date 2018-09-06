package snowflake

import (
	"errors"
	"sync"
	"time"
)

//唯一ID算法
const (
	twepoch      = int64(1483228800000) //开始时间戳（2017-01-01）
	workeridBits = uint(10)             //机器id所占的wuov
	sequenceBits = uint(12)             //序列所占的位数
	workeridMax  = int64(-1 ^ (-1 << workeridBits))
	sequenceMask = int64(-1 ^ (-1 << sequenceBits))

	workeridShift  = sequenceBits
	timestampShift = sequenceBits + workeridBits
)

type Snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

func NewSnowflake(workerid int64) (*Snowflake, error) {
	if workerid < 0 || workerid > workeridMax {
		return nil, errors.New("workerid must be between 0 and 1023")
	}

	return &Snowflake{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

func (s *Snowflake) Generate() int64 {
	s.Lock()
	defer s.Unlock()

	now := time.Now().UnixNano() / 1000000

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}

	s.timestamp = now

	r := int64((now-twepoch)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))
	return r
}
