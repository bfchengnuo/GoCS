package main

import (
	"fmt"
	"time"
)

// 简单的多路复用演示
func main() {
	func2()
}

func func1() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			// 0  2  4  6  8，因为大小为 1
			fmt.Println(x)
		case ch <- i:
			fmt.Println("send ", i)
		default:
			fmt.Println("默认操作")
		}
	}

	time.Sleep(time.Duration(5) * time.Second)
}

// 倒计时模拟
func func2() {
	tick := time.Tick(1 * time.Second)

	for x := 10; x > 0; x-- {
		fmt.Println(x)
		<-tick
	}
}
