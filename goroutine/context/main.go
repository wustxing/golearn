package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	urls := []string{"http://www.baidu.com", "http://www.cnbeta.com"}

	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range urls {
		wg.Add(1)
		subCtx := context.WithValue(ctx, "url", url)
		go reqURL(subCtx, wg)
	}

	go func() {
		time.Sleep(time.Second * 3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value("url").(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop getting url :%s\n", url)
			return
		default:
			r, err := http.Get(url)
			if err == nil && r.StatusCode == http.StatusOK {
				body, _ := ioutil.ReadAll(r.Body)
				fmt.Println("show Body", body[:10])
			}
			r.Body.Close()
			time.Sleep(time.Second * 1)
		}
	}
}
