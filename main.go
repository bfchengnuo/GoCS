package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	// 获取命令行参数，第一个参数为文件名，切片处理
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}
}
