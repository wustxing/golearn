package mutex

import (
	"sync"
)

type Mutex struct {
	sema uint32
}

func NewMutex() *Mutex {
	var m Mutex
	m.sema = 1
	return &m
}

func (m *Mutex) Lock() {
	sync.Runtime_Semrelease(&m.sema)
}
