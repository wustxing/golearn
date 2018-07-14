package main

import "log"

func main() {
	// 先声明map
	var m1 map[string]string
	// 再使用make函数创建一个非nil的map，nil map不能赋值
	m1 = make(map[string]string)
	// 最后给已声明的map赋值
	m1["a"] = "aa"
	m1["b"] = "bb"
	m1["c"] = "cc"
	log.Println(m1)


	// 遍历map
	for k, v := range m1 {
		log.Println(k, v)
		delete(m1,k)
	}


	log.Println(m1)
}
