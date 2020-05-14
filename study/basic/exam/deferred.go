package main

import "fmt"

// defer 关键字测试
func main() {
	// 在 Go 语言的函数中 return 语句在底层并不是原子操作，它分为给返回值赋值和 RET 指令两步。
	// 而 defer 语句执行的时机就在返回值赋值操作后，RET 指令执行前。
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

	println()

	x := 1
	y := 2
	// defer 注册要延迟执行的函数时该函数所有的参数都需要确定其值
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

// 5
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 6
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 5
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
