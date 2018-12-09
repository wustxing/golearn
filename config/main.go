package main

import (
	"fmt"
	"github.com/0990/golearn/config/conf"
	"github.com/0990/golearn/config/manager"
	"log"
	"net/http"
	"time"
)

var confManger *manager.ConfManager
var ch chan func()
var confGorutine *conf.Config

//配置重载demo
func main() {
	confManger = &manager.ConfManager{}
	err := confManger.Init("student_config.pbt")
	if err != nil {
		log.Fatal(err)
	}

	//使用的线程
	ch = make(chan func(), 100)
	//ch <- func() {
	confGorutine = confManger.GetConfig()
	//}
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				conf, _ := confGorutine.GetStudent(1001)
				fmt.Println("before", conf)
				time.Sleep(time.Second * 20)
				conf1, _ := confGorutine.GetStudent(1001)
				fmt.Println("after", conf, conf1)
				//tconf, _ := confManger.GetConfig().GetTeacher(1001)
				//fmt.Println(conf, tconf)
			case f := <-ch:
				f()
			}
		}
	}()

	http.HandleFunc("/reload", handler)
	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params := make(map[string]string)
	for k, v := range r.Form {
		params[k] = v[0]
	}

	command := params["command"]
	fmt.Println("command", command)
	err := confManger.Reload(command)
	if err != nil {
		fmt.Fprint(w, "reload failed")
		return
	}
	ch <- func() {
		confGorutine = confManger.GetConfig()
	}
	fmt.Println("reload conf success")
	fmt.Fprint(w, "reload success")
}
