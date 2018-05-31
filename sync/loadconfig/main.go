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
	fmt.Println("get config enter")
	mutex.RLock()
	fmt.Println("get config begin")
	time.Sleep(3*time.Second)
	fmt.Println("get config end")
	defer mutex.RUnlock()
	return p.name
}


var config *Config
var waitgroup sync.WaitGroup
var mutex sync.RWMutex
func loadConfig(){
	fmt.Println("loadConfig enter")
	mutex.Lock()
	fmt.Println("loadConfig begin")
	defer mutex.Unlock()
	config = &Config{
	}
	time.Sleep(1*time.Second)
	config.name = "xujialong"
	config.id = 2
	fmt.Println("loadConfig end")
	waitgroup.Done()
}
func useConfig(){
	config.GetConfig()
	waitgroup.Done()
}

func init(){
	config = &Config{
		name:"hi",
	}

}

func main(){
	waitgroup.Add(1)
	waitgroup.Add(1)
	go useConfig()
	go loadConfig()
	waitgroup.Wait()
}
