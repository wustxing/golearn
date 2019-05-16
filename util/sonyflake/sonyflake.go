package sonyflake

import (
	"sync"
	"time"
)

const (
	BitLenSequence = 8
)

const sonyflakeTimeUnit = 1e7 // nsec, i.e. 10 msec

type Sonyflake struct {
	mutex       *sync.Mutex
	elapsedTime int64
	sequence    uint16
}

func NewSonyFlake() *Sonyflake {
	sf := new(Sonyflake)
	sf.mutex = new(sync.Mutex)
	//sf.sequence = uint16(1<<BitLenSequence - 1)
	return sf
}

func (sf *Sonyflake) NextID() int64 {
	const maskSequence = uint16(1<<BitLenSequence - 1)
	sf.mutex.Lock()
	defer sf.mutex.Unlock()

	current := time.Now().UnixNano() / sonyflakeTimeUnit

	if sf.elapsedTime < current {
		sf.elapsedTime = current
		sf.sequence = 0
	} else {
		sf.sequence = (sf.sequence + 1) & maskSequence
		if sf.sequence == 0 {
			sf.elapsedTime++
			overTime := sf.elapsedTime - current
			time.Sleep(sleepTime(overTime))
		}
	}
	return sf.toID()
}

func sleepTime(overtime int64) time.Duration {
	return time.Duration(overtime)*10*time.Millisecond -
		time.Duration(time.Now().UTC().UnixNano()%sonyflakeTimeUnit)*time.Nanosecond
}

func (sf *Sonyflake) toID() int64 {
	return int64(sf.elapsedTime)<<(BitLenSequence) | int64(sf.sequence)
}
