package main

import (
	"fmt"
	"net/url"
)

// 解析相关
func main() {
	// url 转义
	data := url.QueryEscape("k1=v1&k2=v 2")
	fmt.Println(data)
}
