package main

import "fmt"

func main() {
	// 定义
	var ch1 chan int
	// 第二个参数是缓冲区大小，可选
	var ch2 = make(chan int, 20)

	fmt.Println(ch1 == nil) // 零值，必须使用 make 初始化后才能使用

	// send
	ch2 <- 233
	ch2 <- 2233
	// receive
	fmt.Println(<-ch2)
	// 丢弃
	<-ch2

	// 手动关闭，非必须，可由 GC 感知回收
	close(ch2)
}

// 使用 channel 串联 goroutine
func series() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 使用匿名函数
	go func() {
		for i := 1; i < 11; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		// 使用 range 接收，相比直接 <-
		// 当通道没有值后，会直接结束，不会一直返回零值
		for x := range ch1 {
			// 运算
			ch2 <- x * x
		}
		close(ch2)
	}()

	for x := range ch2 {
		fmt.Println(x)
	}
}

// 单向通道
// chan<- int  单向向通道输出
// x <- chan int 单向从通道输出
// 双向可以变单向，反之不可
func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

var done = make(chan struct{})

// 通过 close chan 并发退出
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
