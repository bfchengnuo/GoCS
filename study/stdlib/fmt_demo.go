package main

import (
	"errors"
	"fmt"
	"os"
)

// fmt 的 print 系列基本每一个都是三个组成，分别对应普通、格式化（xxf）、换行（xxxln）
func main() {
	//basic()
	//scanFunc()
	scanFunc2()
}

func basic() {
	// 向 io.Writer 输出
	// 例如标准输出流，或者文件也可以
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")

	// 将结果以字符串的方式返回，不是打印！
	str := fmt.Sprint("返回字符串")
	fmt.Println(str)

	// 定义错误
	err := fmt.Errorf("这是一个错误")
	e := errors.New("原始错误e")
	ew := fmt.Errorf("wrap 了一个错误 %w", e)
	fmt.Printf("err %T - %s \n", err, err)
	fmt.Println(ew)

	// 宽度标识符
	fmt.Printf("num: %4.3f", 23.1235)
}

func scanFunc() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scan(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)

	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}

func scanFunc2() {
	// 从 io.Reader 中读取数据
	var name string
	var desc string
	fmt.Fscan(os.Stdin, &name)
	fmt.Println("name: ", name)

	fmt.Sscan("str...", &desc)
	fmt.Println("desc: ", desc)
}
