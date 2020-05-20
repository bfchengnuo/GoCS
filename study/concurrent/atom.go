package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作
func main() {
	// 并发安全且比互斥锁效率更高
	c3 := AtomicCounter{}
	test(&c3)
}

type Counter interface {
	Inc()
	Load() int64
}

type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}
