package pool

import (
	"io"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	maxGoroutines   = 5
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (p *dbConnection) Close() error {
	log.Println("Close:Connection", p.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create new connection", id)
	return &dbConnection{ID: id}, nil
}

func Test_Pool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines * 2)

	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Fatal(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		time.Sleep(time.Millisecond * 100)
		go func(q int) {
			performQueries(query, p)
			wg.Done()
		}(query)
	}
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(query, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("ShutDown Program")
	p.Close()
}

func performQueries(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)

	//time.Sleep(10 * time.Millisecond)
	log.Printf("query:%d conn:%d", query, conn.(*dbConnection).ID)
}
