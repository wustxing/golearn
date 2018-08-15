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

func redisSet(key string, value string) {
	c, err := RedisClient.Dial()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	_, err = c.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

func redisGet(key string) (value string) {
	c, err := RedisClient.Dial()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	val, err := redis.String(c.Do("GET", key))

	if err != nil {
		fmt.Println("redis get failed:", err)
		return ""
	} else {
		fmt.Printf("Got value is %v\n", val)
		return val
	}
}

func redisIncr(key string) (value string) {
	c, err := RedisClient.Dial()
	_, err = c.Do("INCR", key)

	if err != nil {
		fmt.Println("incr error", err.Error())
	}

	incr, err := redis.String(c.Do("GET", key))

	if err != nil {
		fmt.Println("redis key after incr is:", incr)
	}
	return incr
}

func main() {
	initRedisPool()
	redisSet("hello", "1")
	value := redisGet("hello")

	redisIncr("hello")
	newValue := redisGet("hello")
	fmt.Println(value, newValue)
}
