package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	// 定义常量
	RedisClient *redis.Pool
)

const (
	RedisAddress = "127.0.0.1:6379"
	RedisDb      = 0
)

func initRedisPool() {
	RedisClient = &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisAddress)
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", RedisDb)
			return c, nil
		},
	}
}

func main() {
	initRedisPool()
	conn := RedisClient.Get()
	defer conn.Close()
	vs, err := redis.Int(conn.Do("HGET", "area_0:invite:userid2count", 1000000001))
	if err != nil {
		if err == redis.ErrNil {
			fmt.Println(err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println(vs)
}
