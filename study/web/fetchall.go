package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 并发版的 fetch
func main() {
	url := "http:baidu.com"
	start := time.Now()
	ch := make(chan string)

	// should for...
	// go 表示创建一个新对 goroutine, 异步执行
	go fetch(url, ch)
	// should for...
	fmt.Println(<-ch) // 阻塞
	fmt.Printf("%.2fs done.\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send, 阻塞
		return
	}
	// 丢弃内容
	bBytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("reading %s: %v", url, err) // send, 阻塞
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, bBytes, url)
}
