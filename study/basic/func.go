package main

import "fmt"

// 函数相关
func main() {
	// 返回值测试
	fmt.Println(func1())

	println()

	// 闭包示例
	var f = adder()
	fmt.Println(f(10)) // 10
	fmt.Println(f(20)) // 20 + 10 = 30
	f1 := adder()
	fmt.Println(f1(30)) // 30
	fmt.Println(f(10))  // 40

	println()

	// 测试 defer
	// 更多信息参考：./exam/deferred.go
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

// 高阶函数，接收一个函数或者返回一个函数（闭包）
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

// 可变参数
func sum(nums ...int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

// 交换变量
func swap(a, b *int) {
	*a, *b = *b, *a
}

// 多返回的两种方式
func div(a, b int) (q, r int) {
	return a / b, a % b
}
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	// 等价于 return q, r
	return
}

// 闭包
// 函数 + 引用环境 = 闭包，返回的函数引用了外部变量 x，并且会一直持有
func adder() func(int) int {
	var x int
	return func(y int) int {
		// x 保存在函数内，不同 adder 创建的互不干扰
		x += y
		return x
	}
}

// 返回值测试，3
func func1() (v int) {
	v = 1
	return 3
}

// 异常处理
func crashFunc() {
	defer func() {
		// 如果程序出出现了 panic 错误,可以通过 recover 恢复过来
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}
