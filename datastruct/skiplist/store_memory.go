package skiplist

import "sync"

type memoryStore struct {
	name string
	m    map[string][]byte
	sync.RWMutex
}

func newMemoryScore(name string) *memoryStore {
	return &memoryStore{name: name, m: make(map[string][]byte)}
}

func (p *memoryStore) getKV(key []byte) ([]byte, error) {
	p.RLock()
	defer p.RUnlock()

	data, ok := p.m[string(key)]
	if !ok {
		return nil, errNotFound
	}
	return data, nil
}

//原子操作，更新及删除
func (p *memoryStore) commit(updates map[string][]byte, dels [][]byte, delRange [][2][]byte) error {
	p.Lock()
	defer p.Unlock()

	for _, v := range dels {
		delete(p.m, string(v))
	}

	for k, v := range updates {
		p.m[k] = v
	}

	//TODO
	//for _, _ := range delRange {
	//
	//}

	return nil
}

func (p *memoryStore) close() error {
	return nil
}
