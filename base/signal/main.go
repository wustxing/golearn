package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

//SIGKILL SIGSTOP无法被捕捉
//SIGINT SIGTERM SIGHUP默认处理是关闭程序
//SIGQUIT 默认处理会打印dump后退出程序
//SIGSTOP和SIGCONT组合使用可以暂停和恢复程序
//示例中setupsigusr1Trap 当kill -SIGUSR1 pid时,会打印dump（不会退出）,用于查日志
//更多 https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter16/16.03.html?h=signal
//Kill 默认发的是syscall.SIGTERM信号

func main() {
	setupSigusr1Trap()
	go print()
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	s := <-c
	fmt.Println("Got signal:", s)
}

func print() {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			fmt.Println("ticker come,", time.Now())
		}
	}

}

func setupSigusr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			DumpStacks()
		}
	}()
}

func DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("===BEGIN dump===\n%s\n===END dump", buf)
}
