package main

import "fmt"

// 常见内建函数
func main() {

}

// new 是内建函数，并非关键字，所以可以修改其定义
// new 函数会返回地址，类似语法糖，可以创建匿名变量，本质与直接定义无区别
// 也许创建结构体等复杂类型的时候用的多一点
// 谨慎使用大小为 0 的类型，避免 GC 的迷惑行为（struct{}、[0]int）
func newFunc() {
	num := new(int)
	fmt.Println(num)
	fmt.Println(*num)
}

// 创建一个指定的匿名数据类型
// make 只用于 slice、map 以及 channel 的初始化，返回的是引用本身
// 为了避免 panic，这三种类型尽量使用 make 创建
func makeFun() {
	// 创建切片，参数为：类型、len、cap，cap 可以省略，默认与 len 相同
	sp := make([]string, 5, 10)
	// 创建 map
	dict := make(map[int]string)

	fmt.Println(sp)
	fmt.Println(dict)
}
