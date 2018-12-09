package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
)

/*set并集 交集 差集计算示例*/

func main() {
	a := set.New(set.ThreadSafe)
	a.Add(1)
	a.Add(2)
	a.Add(3)

	b := set.New(set.ThreadSafe)
	b.Add(2)
	b.Add(3)
	b.Add(4)

	//并集
	unionSet := set.Union(a, b)
	fmt.Printf("并集:%v\n", unionSet)

	//交集
	intersectionSet := set.Intersection(a, b)
	fmt.Printf("交集:%v\n", intersectionSet)

	//差集
	diffS1S2 := set.Difference(a, b)
	fmt.Printf("差集(属a不属b):%v\n", diffS1S2)

	diffS2S1 := set.Difference(b, a)
	fmt.Printf("差集(属b不属a):%v\n", diffS2S1)
}
