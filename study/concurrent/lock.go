package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var lwg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func main() {
	mutexFunc()
}

// 互斥锁
// 使用最广泛的锁
func mutexFunc() {
	lwg.Add(2)
	go add()
	go add()
	lwg.Wait()
	fmt.Println(x)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	lwg.Done()
}

// 读写锁
func write() {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwlock.Unlock() // 解写锁
	lwg.Done()
}

func read() {
	rwlock.RLock() // 加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
	lwg.Done()
}
