package main

import "fmt"

// 作用域示例
func main() {
	x := "hello"
	for _, x := range x {
		x := string(x) + "- inner"
		fmt.Println(x)
	}
	// if 也是同理，注意词法域
}

var name = "mps"

func scopeTest() {
	fmt.Println("out name: ", name)
	name := "LH"
	fmt.Println("func name: ", name)
}
