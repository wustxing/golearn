package main

import (
	"sync"
	"fmt"
	"time"
)

type safe_set struct{
	sync.Mutex
	s []interface{}
}

func (p *safe_set)Iter() chan interface{}{
	ch :=make(chan interface{},len(p.s))

	go func(){
		p.Lock()
		defer p.Unlock()
		for e,v:=range p.s{
			fmt.Println("insert start",e)
			time.Sleep(time.Second*2)
			fmt.Println("insert end",e)
			ch<-e
			fmt.Println(e,v)
		}
		time.Sleep(time.Second*5)
		close(ch)
	}()
	return ch
}
func main() {
	th := safe_set{
		s: []interface{}{"a", "b"},
	}
	ch:=th.Iter()
	v:=<-ch
	//v := <-th.Iter()

	//fmt.Sprintf("%s%v", "ch", v)
	fmt.Println(v)
	time.Sleep(time.Second*10)
}