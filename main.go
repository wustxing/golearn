package main

import "fmt"

type hello struct{
	intValue *int
}

func (p *hello)IntValue() int{
	return *(p.intValue)
}
func main() {
	intArr :=make([]*hello,0)
	a:=1
	b:=2
	intArr  = append(intArr,&hello{
		intValue:&a,
	})
	intArr = append(intArr,&hello{
		intValue:&b,
	})
	hello:=getHello(1,intArr)
	fmt.Println(hello.IntValue())

}

func getHello(intValue int,intArr []*hello)*hello{
	for _,v:=range intArr{
		if v.IntValue() ==intValue{
			//return intArr[i]
			return v
		}
	}
	return nil
}
