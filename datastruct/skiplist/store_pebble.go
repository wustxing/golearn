package skiplist

import (
	"github.com/cockroachdb/pebble"
)

type pebbleStore struct {
	name string
	db   *pebble.DB
}

func newPebbleScore(name string) (*pebbleStore, error) {
	db, err := pebble.Open(name, &pebble.Options{
		EventListener: pebble.MakeLoggingEventListener(pebble.DefaultLogger),
	})
	if err != nil {
		return nil, err
	}

	return &pebbleStore{name: name, db: db}, nil
}

func (p *pebbleStore) getKV(key []byte) ([]byte, error) {
	data, closer, err := p.db.Get(key)
	if err != nil {
		if err == pebble.ErrNotFound {
			return nil, errNotFound
		}
		return nil, err
	}

	defer closer.Close()

	ret := make([]byte, len(data))
	copy(ret, data)

	return ret, nil
}

//原子操作，更新及删除
func (p *pebbleStore) commit(updates map[string][]byte, dels [][]byte, delRange [][2][]byte) error {
	b := p.db.NewBatch()
	defer b.Close()

	for _, v := range dels {
		err := b.Delete(v, &pebble.WriteOptions{})
		if err != nil {
			return err
		}
	}

	for k, v := range updates {
		err := b.Set([]byte(k), v, &pebble.WriteOptions{})
		if err != nil {
			return err
		}
	}

	for _, v := range delRange {
		b.DeleteRange(v[0], v[1], &pebble.WriteOptions{})
	}

	return b.Commit(pebble.NoSync)
}

func (p *pebbleStore) close() error {
	return p.db.Close()
}
