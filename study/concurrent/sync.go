package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 内存同步测试
// 内存可见性，多个核心，各自有自己的缓存，并且是互相不可见
// CPU 的缓冲区，批量处理完毕后再 flush 到主存
func main() {
	var x, y int

	go func() {
		x = 1
		fmt.Println("y = ", y)
	}()

	go func() {
		y = 1
		fmt.Println("x = ", x)
	}()

	// 可能出现 0 0 的结果
	time.Sleep(1 * time.Second)
}

// 并发安全的 map
var m = sync.Map{}

func mapFunc() {
	for i := 0; i < 10; i++ {
		go func(n int) {
			key := strconv.Itoa(n)
			// save
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
		}(i)
	}
}
