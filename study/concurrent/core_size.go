package main

import (
	"fmt"
	"runtime"
	"time"
)

// Go1.5版本之前，默认使用的是单核心执行。Go1.5版本之后，默认使用全部的CPU逻辑核心数。
// 参考 GPM 调度模型，Go 的 goroutine 调度是自己实现，在用户态完成，不涉及 OS 调度，速度自然快
func main() {
	// 设置单核心(逻辑)，相当于串行
	// m:n 模型，将 m 个 goroutine 调度到 n 个 os 线程，这里设置的是 n
	runtime.GOMAXPROCS(1)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
