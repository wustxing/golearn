package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const (
	rankCopyScript = `
        local rankData = redis.call("ZREVRANGE",KEYS[1],0,ARGV[1]-1,"WITHSCORES")
		if rankData then
			redis.call("DEL",KEYS[2])
			--local step = 1
        	for i = 1, #rankData, 2 do
        		redis.call("ZADD",KEYS[2], rankData[i+1],rankData[i])
        	end
		end
		return rankData[3]
	`
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	//insertBefore := time.Now()
	//for i := 0; i < 100000; i++ {
	//	client.ZAdd("rank", redis.Z{
	//		Score:  float64(i),
	//		Member: i + 1,
	//	})
	//}
	//fmt.Println("insert elaspse", time.Since(insertBefore))

	//rankBefore := time.Now()
	//result, err := client.ZRevRangeWithScores("rank_test_12", 0, 10000).Result()
	//fmt.Println("rank get  elaspse", time.Since(rankBefore), err)
	//_, err = client.ZAdd("rank_test_copy", result...).Result()
	//fmt.Println("rank copy elaspse", time.Since(rankBefore), err)
	//rankCopyBefore := time.Now()
	//client.ZInterStore("rank_test_copy", redis.ZStore{}, "rank_test_12")
	//fmt.Println("rank copy elaspse", time.Since(rankCopyBefore))

	//client.Z
	copyBefore := time.Now()
	ret, err := client.Eval(rankCopyScript, []string{"rank", "rank_copy"}, 10000).Result()
	fmt.Println(ret, err)
	fmt.Println("copy elaspse", time.Since(copyBefore))

}
