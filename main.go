package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//hash := HashUserID(1000067)
	//orderID := CreateOrderID(hash)
	//hashID := orderID[12:13]
	//fmt.Println(orderID, hash, hashID)
	//fmt.Println(time.Now().Local().Format("20060102150405"))
	fmt.Println(CreateOrderIDNew(114))
}

func CreateOrderID(hashID int32) string {
	return fmt.Sprintf("%s%d%d", time.Now().Local().Format("200601021504"), hashID, rand.Int63n(90000000)+10000000)
}

func HashUserID(userID uint64) int32 {
	return int32(userID % 10)
}

func CreateOrderIDNew(serverid int32) string {
	return fmt.Sprintf("%d%s%d", serverid, time.Now().Local().Format("200601021504"), rand.Int63n(9000000)+1000000)
}
