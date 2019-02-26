package util

import (
	"errors"
	"fmt"
	"math"
	"sanguosha.com/sgs_nuyan/gameutil"
	"sync"
	"time"
)

const sequenceMax = 1e5 // 2018/1/1 00:00:00

type IDManager struct {
	serverID  uint32
	maxSeqID  uint64
	lastSeqID uint64
	timestamp int64
	mutex     sync.Mutex
}

func (p *IDManager) Init(serverID uint32) error {
	maxServerID := uint32(math.Pow(2, 12))
	if serverID > maxServerID {
		return errors.New("serverID too max")
	}
	p.serverID = serverID
	p.maxSeqID = uint64(math.Pow(10, 3))
	return nil
}

// GeneratePKID ...
func (p *IDManager) GeneratePKID() uint64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	now := gameutil.GetCurrentTimestamp()
	if p.timestamp == now {
		p.lastSeqID++
		if p.lastSeqID >= sequenceMax {
			p.lastSeqID = 0
			for now <= p.timestamp {
				time.Sleep(time.Second)
				fmt.Print("over", p.lastSeqID)
				now = gameutil.GetCurrentTimestamp()
			}
		}
	} else {
		p.lastSeqID = 0
	}
	p.timestamp = now
	return uint64(now)*1e9 + uint64(p.serverID)*sequenceMax + p.lastSeqID
}
