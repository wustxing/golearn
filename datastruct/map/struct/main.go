package main

import "fmt"

type MatchUserID int32
type Key struct{
	id1  MatchUserID
	id2  MatchUserID
}

func main(){
	testMap:=make(map[Key]int)
	testMap[Key{1,2}] = 20

	testMap[Key{2,1}] = 300

	testMap[Key{2,1}] = 400

	fmt.Println(testMap)
}
