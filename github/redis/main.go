package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

/*redis库使用示例*/
//windows redis下载地址： https://github.com/MicrosoftArchive/redis/releases

func main() {
	//创建一个redispool池
	rp := &redis.Pool{
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	//从中取出一个连接
	conn := rp.Get()
	defer conn.Close()

	////创建key为"name"值为"0990"的键值对
	//_, err := conn.Do("SET", "name", "0990")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////取出key为"name"的值并打印
	//val, err := redis.String(conn.Do("GET", "name"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(val)

	//vs, err := redis.Values(conn.Do("ZREVRANGE", "area_194:rank:level", 0, 100, "WITHSCORES"))
	//if err != nil {
	//	logrus.Debug(vs)
	//}
	//ivs := make([]string, len(vs))
	//err = redis.ScanSlice(vs, &ivs)
	//if err != nil {
	//	logrus.Debug(err)
	//}
	//fmt.Println(ivs)
	//
	//ret, err := redis.Strings(conn.Do("ZREVRANGE", "area_194:rank:level", 0, 100))
	//fmt.Println(ret)

	conn.Do("HMSET", "xujialong", "key1", "value1", "key2", "value2")
}
