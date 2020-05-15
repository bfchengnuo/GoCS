package main

import (
	"fmt"
	"sort"
)

// 集合操作相关
func main() {
	sp := []string{"mps", "skye", "yh", "yousa", "kerronex"}
	// 排序
	sort.Strings(sp)
	fmt.Println(sp)

	// map
	dict := map[int]string{1: "ONE"}
	val, ok := dict[2]
	if !ok {
		fmt.Println("key 不存在")
		return
	}
	fmt.Println("map-val: ", val)
}
