package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Panic(err)
	}

	if _, err := conn.Do("AUTH", "Xu2009"); err != nil {
		conn.Close()
		log.Panic(err)
	}

	type Score struct {
		x, y int32
		v    int64
	}

	rand.Seed(time.Now().UnixNano())

	m := make(map[int32]Score)
	for i := 0; i < 20; i++ {
		x := rand.Int31() % 100000
		y := rand.Int31() % 64
		m[x] = Score{
			x: x,
			y: y,
			v: Marshal(x, y),
		}
	}

	for _, v := range m {
		conn.Do("ZADD", "ranktest", v.v, v.x)
	}

	ret, err := redis.ByteSlices(conn.Do("ZRANGE", "ranktest", 0, -1, "withscores"))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(ret)

	output := make(map[int32]Score)
	for i := 0; i < len(ret); i += 2 {
		x, errx := strconv.ParseInt(string(ret[i]), 10, 64)
		v, errv := strconv.ParseFloat(string(ret[i+1]), 64)
		fmt.Println(x, int64(v), errx, errv)

		score, extend := Unmarshal(int64(ret[i+1]))
		output[int32(x)] = Score{
			x: score,
			y: extend,
			v: int64(v),
		}
	}

	for _, v := range output {
		v1, ok := m[v.x]
		if !ok {
			fmt.Println("key not exist", v.x)
		}

		fmt.Println(v, v1)
	}

}

func Marshal(score int32, extend int32) int64 {
	now := time.Now().Unix()
	return (int64(score&0xffffff) << 40) | ((now & 0x3ffffffff) << 6) | int64(extend&0x3f)
}

func Unmarshal(s int64) (score int32, extend int32) {
	extend = int32(s & 0x3f)
	t := (s >> 6) & 0x3ffffffff
	score = int32((s >> 40) & 0xffffff)

	fmt.Println(t)
	return score, extend
}
