package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Go 的并发主要依赖 goroutine 和 channel
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)   // 计数 +1，要在 goroutine 创建之前调用
		go hello(i) // 值传递，copy
	}

	// 等待所有 wg 完成
	wg.Wait()
}

func hello(i int) {
	defer wg.Done() // goroutine 结束就登记 -1
	fmt.Println("Hello Goroutine!", i)
}
