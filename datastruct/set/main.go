package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
	"sort"
)

func main() {
	s1 := set.New()
	s1.Add(1)
	s1.Add(2)
	s1.Add(2)
	s1.Add(5)
	s1.Add(4)

	s2 := set.New()
	s2.Add(2)
	s2.Add(3)

	diffs1 := set.Difference(s1, s2)
	s3 := set.Union(s1, s2)
	//diffs2 := set.Difference(s2, s1)
	//fmt.Println(s1, s2, diffs1, diffs2)
	//sort.Sort(diffs1.List())
	list := diffs1.List()
	var intlist []int
	for _, v := range list {
		intlist = append(intlist, v.(int))
	}
	sort.Slice(intlist, func(i, j int) bool {
		return intlist[i] < intlist[j]
	})
	fmt.Println(diffs1, intlist, s3)
	//fmt.Println(s3)
	//for _, v := range s1.List() {
	//	intValue := v.(int)
	//	fmt.Println(intValue)
	//}
}
