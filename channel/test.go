package main

import (
	"sync"
	"fmt"
	"reflect"
)

type safe_set struct{
	sync.Mutex
	s []interface{}
}

func (p *safe_set)Iter() chan interface{}{
	ch :=make(chan interface{})

	go func(){
		p.Lock()
		defer p.Unlock()
		for e,v:=range p.s{
			ch<-e
			fmt.Println(e,v)
		}
		close(ch)
	}()
	return ch
}
func main() {
	th := safe_set{
		s: []interface{}{"a", "b"},
	}
	v:=<-th.Iter()
	fmt.Println(reflect.TypeOf(v))
	fmt.Sprintf("%s%v","ch",v)
	v=<-th.Iter()
	fmt.Println(reflect.TypeOf(v))
	fmt.Sprintf("%s%v","ch",v)
}