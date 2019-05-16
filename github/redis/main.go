package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sanguosha.com/sgs_nuyan/gameutil"
	"sanguosha.com/sgs_nuyan/proto/db"
	"strconv"
	"time"
)

/*redis库使用示例*/
//windows redis下载地址： https://github.com/MicrosoftArchive/redis/releases
//const (
//	rankCopyScript = `
//        local rankData = redis.call("ZREVRANGE",KEYS[1],0,ARGV[1]-1,"WITHSCORES")
//		if rankData then
//			redis.call("DEL",KEYS[2])
//			local step = 3000
//        	for i = 1, #rankData, step do
//        		redis.call("ZADD",KEYS[2], unpack(rankData, i, math.min(i + step - 1, #rankData)))
//        	end
//		end
//        local retData = redis.call("ZREVRANGE",KEYS[2],0,100,"WITHSCORES")
//		return retData
//	`
//
//	myRank = `
//		local rank = redis.call("ZREVRANK",KEYS[1],ARGV[1])
//		if rank then
//			local rankInt = tonumber(rank)
//			local retData = redis.call("ZREVRANGE",KEYS[1],math.max(rankInt-1,0),rankInt+1,"WITHSCORES")
//			return {retData,rank}
//		end
//		return rank
//	`
//)

const (
	rankCopyScript = `
        local rankData = redis.call("ZREVRANGE",KEYS[1],0,ARGV[1]-1,"WITHSCORES")
		if rankData then
			redis.call("DEL",KEYS[2])
        	for i = 1, #rankData, 2 do
        		redis.call("ZADD",KEYS[2], rankData[i+1],rankData[i])
        	end
		end
        local retData = redis.call("ZREVRANGE",KEYS[2],0,ARGV[2]-1,"WITHSCORES")
		return retData
	`
	ScriptGetQualifyRankInfo = `
        local rankData = redis.call("ZREVRANGE",KEYS[1],0,ARGV[1]-1,"WITHSCORES")
		if rankData then
			local rank = redis.call("ZREVRANK",KEYS[1],ARGV[2])
			if rank then
				local rankInt = tonumber(rank)
				local myRankData = redis.call("ZREVRANGE",KEYS[1],math.max(rankInt-1,0),rankInt+1,"WITHSCORES")
				return {rankData,rank,myRankData}
			end
			return {rankData,0,0}
		end
		return rankData
	`
)

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

	//script := redis.NewScript(2, rankCopyScript)
	//ret, err := redis.Strings(script.Do(conn, "hello", "helloCopy", 10000, 100))
	//if err != nil {
	//	logrus.WithError(err).Error("Make Rank error")
	//	return
	//}
	//fmt.Print(ret)

	//reply, err := conn.Do("ZREVRANGE", "rank", 0, 10000, "WITHSCORES")
	//if err != nil {
	//	return
	//}
	//conn.Do("ZADD", "rank_copy")

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

	////conn.Do("HMSET", "xujialong", "key1", "value1", "key2", "value2")
	//copyBefore := time.Now()
	//script := redis.NewScript(1, myRank)
	//rawReply, _ := redis.Values(script.Do(conn, "rank", 99998))
	//fmt.Println(rawReply)
	//reply, err := redis.Strings(rawReply[0], err)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(reply)
	//myRank, err := redis.Int64(rawReply[1], err)
	//fmt.Println(myRank)
	//for i := 0; i < len(reply); i += 2 {
	//	//int1, _ := strconv.ParseUint(reply[i], 10, 64)
	//	//int2, _ := strconv.ParseUint(reply[i+1], 10, 64)
	//	//fmt.Println(int1, int2)
	//}
	//fmt.Println("copy elaspse", time.Since(copyBefore))
	//rank, err := redis.Strings(conn.Do("ZREVRANGE", "rank", 1, 2, "WITHSCORES"))
	//if err == redis.ErrNil {
	//	fmt.Println("not existed")
	//}
	//fmt.Println(rank, err)
	userID := uint64(30000041030)
	recordAll, myRank, recordMy, err := GetQualifyRankInfo(conn, userID)
	fmt.Println(recordAll)
	fmt.Println(myRank)
	fmt.Println(recordMy)
	fmt.Println(err)
}

func GetQualifyRankInfo(conn redis.Conn, userID uint64) (*db.RankRecord, int32, *db.RankRecord, error) {
	rankRecordMy := &db.RankRecord{}

	script := redis.NewScript(1, ScriptGetQualifyRankInfo)
	rawReply, err := redis.Values(script.Do(conn, "area_194:qualify_19031501:rank", 100, gameutil.UserKey(userID)))
	if err != nil {
		return nil, 0, rankRecordMy, err
	}

	ivs, err := redis.Strings(rawReply[0], err)
	if err != nil {
		return nil, 0, rankRecordMy, err
	}

	rankRecord := &db.RankRecord{}
	for i := 0; i < len(ivs); i += 2 {
		info := &db.RankRecord_Item{}
		info.Userid, err = strconv.ParseUint(ivs[i], 10, 64)
		info.Score, err = strconv.ParseInt(ivs[i+1], 10, 64)
		rankRecord.List = append(rankRecord.List, info)
	}

	myRank, myRankErr := redis.Int64(rawReply[1], err)
	if myRankErr != nil {
		return rankRecord, 0, rankRecordMy, myRankErr
	}

	if myRank > 0 {
		ivs, err = redis.Strings(rawReply[2], err)
		if err != nil {
			return rankRecord, int32(myRank + 1), rankRecordMy, err
		}

		for i := 0; i < len(ivs); i += 2 {
			info := &db.RankRecord_Item{}
			info.Userid, err = strconv.ParseUint(ivs[i], 10, 64)
			info.Score, err = strconv.ParseInt(ivs[i+1], 10, 64)
			rankRecordMy.List = append(rankRecordMy.List, info)
		}
		return rankRecord, int32(myRank + 1), rankRecordMy, err
	} else {
		return rankRecord, 0, rankRecordMy, err
	}

}
