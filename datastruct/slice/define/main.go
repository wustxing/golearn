package define

import "fmt"

func main(){
	//空切片
	slice0:=make([]int,0)
	slice0=append(slice0,1)
	slice2:=[]int{}
	slice2 = append(slice2,1)

	//nil切片
	var slice1 []int
	slice1 = append(slice1,1)
	fmt.Println(slice0,slice1,slice2)
}