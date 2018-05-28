package main

import (
	"time"
	"fmt"
	"sync"
)

type Config struct{
	name string
	id int
}

func(p *Config)GetConfig()string{
	mutex.RLock()
	defer mutex.RUnlock()
	return p.name
}


var config *Config
var waitgroup sync.WaitGroup
var mutex sync.RWMutex
func loadConfig(){
	mutex.Lock()
	defer mutex.Unlock()
	config = &Config{
	}
	time.Sleep(2*time.Second)
	config.name = "xujialong"
	config.id = 2
	println("load over")
	waitgroup.Done()
}
func useConfig(){
	time.Sleep(1*time.Second)
	fmt.Println(config.GetConfig())
	fmt.Println("hello")
	waitgroup.Done()
}

func init(){
	config = &Config{
		name:"hi",
	}

}

func main(){
	waitgroup.Add(1)
	go loadConfig()
	waitgroup.Add(1)
	go useConfig()
	waitgroup.Add(1)
	go useConfig()
	waitgroup.Wait()
}
