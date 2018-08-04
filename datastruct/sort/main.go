package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	Id   int
}

//type Persons []*Person

func (p PersonList) Len() int {
	return len(p)
}

func (p PersonList) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// 交换数据
func (p PersonList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	ages := []int{2, 1, 5, 66, 55, 23, 78, 98, 13}
	sort.Ints(ages)
	for _, value := range ages {
		fmt.Println(value)
	}

	personMap := make(map[int]*Person)
	personMap[1] = &Person{
		Name: "test1",
		Age:  10,
		Id:   1,
	}

	personMap[4] = &Person{
		Name: "test2",
		Age:  40,
		Id:   4,
	}

	personMap[3] = &Person{
		Name: "test2",
		Age:  30,
		Id:   3,
	}

	fmt.Println("排序前")
	for k, person := range personMap {
		fmt.Println(k, person)
	}

	fmt.Println("排序后")
	sortList := sortMapByValue(personMap)
	for _, person := range sortList {
		fmt.Println(person)
	}

	for k, person := range personMap {
		fmt.Println(k, person)
	}
}

type PersonList []*Person

func sortMapByValue(m map[int]*Person) PersonList {
	p := make(PersonList, 0)

	//i := 0

	for _, v := range m {
		p = append(p, v)
	}

	sort.Sort(p)
	return p
}
