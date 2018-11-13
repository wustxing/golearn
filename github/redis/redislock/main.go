package main

import (
	"fmt"
	"github.com/everalbum/redislock"
	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
	"log"
	"time"
)

const (
	RedisAddress = "127.0.0.1:6379"
	RedisDb      = 0
)

func main() {
	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisAddress)
		},
		MaxIdle:   0,
		MaxActive: 3,
		Wait:      true,
	}

	conn := redisPool.Get()
	lock, err := TryLock(conn, "xujialong")
	if err != nil {
		log.Fatal("error while attempting lock")
	}
	defer lock.Unlock()
	fmt.Println("a getlock")
	time.Sleep(20 * time.Second)
}

func TryLock(conn redis.Conn, resource string) (lock *redislock.Lock, err error) {
	for {
		lock, ok, err := redislock.TryLockWithTimeout(conn, resource, time.Second*30)
		if err != nil {
			return nil, errors.New("error while attempting lock")
		}
		if !ok {
			time.Sleep(time.Second)
			continue
		}
		return lock, nil
	}
}
