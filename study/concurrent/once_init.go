package main

import (
	"fmt"
	"sync"
)

var loadOnce sync.Once
var config map[string]string

// 一次初始化
func main() {
	getConfig()
}

func getConfig() map[string]string {
	// 只有一个 do 方法，如果需要传递参数，请使用闭包方式
	// 类似于保存一个 boolean 变量，第一次未初始化是 false，使用互斥锁初始化，之后变为 true，直接返回
	loadOnce.Do(loadConfig)
	// 然后可以返回一个全局变量
	return config
}

func loadConfig() {
	config = make(map[string]string)
	fmt.Println("这里可以初始化全局变量")
}

// 单例模式
type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
