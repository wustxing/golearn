package main

import "fmt"

func main(){
	//由于newSlice和slice 共用空间，给newSlice附加20,也改变了原slice值
	slice:=[]int{1,2,3,4,5}
	newSlice:=slice[1:3]
	fmt.Println(slice,newSlice)
	newSlice=append(newSlice, 20)
	fmt.Println(slice,newSlice)

	//超过长度，会复制一个新的数组
	newSlice1:=slice[4:5]
	newSlice1 = append(newSlice1,30)
	fmt.Println(slice,newSlice,newSlice1)
	//改变切片数组里的值也不影响，因为这是新的数组
	newSlice1[0] = 40
	fmt.Println(slice,newSlice,newSlice1)

	newSlice2:=append(slice,50)
	newSlice3:=append(slice,slice...)
	fmt.Println(slice,newSlice,newSlice1,newSlice2,newSlice3)

	//为了避免影响原切片，可以复制出一个切片
	copySlice:=make([]int,3)
	copy(copySlice,slice[1:3])
	copySlice[0]=60
	fmt.Println(slice,copySlice)
}
