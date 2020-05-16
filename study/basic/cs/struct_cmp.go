package main

import "fmt"

// 结构体地址测试代码
func main() {
	s1 := str{"1"}
	s2 := str{"1"}
	fmt.Println(s1 == s2) // true, 结构体内部可比较
	fmt.Printf("s1: %p\n", &s1)
	fmt.Printf("s2: %p\n", &s2)
	println(s1.txt)

	println()

	s3 := &str{"2"}
	s4 := &str{"2"}
	fmt.Println(s3 == s4) // false
	println(s3.txt)
}

type str struct {
	txt string
}
