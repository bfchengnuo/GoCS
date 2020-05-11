package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// 字符串相关测试
// 字符串处理的相关 API 在 strings 包中
// 多行字符串使用反引号 `` 定义
func main() {
	str := "abc 中英文混合字符串测试"

	for i, v := range str {
		// 坐标是字节，v 是对应的 u8 编码值
		fmt.Println(i, string(v))
	}

	// 仅用于测试的时候用，可能会出现奇怪的问题，builtin 包下
	println()

	// 正确的坐标遍历
	for i, c := range []rune(str) {
		fmt.Println(i, string(c))
	}

	println()

	// 统计长度
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))

	println()

	// 切片
	fmt.Println(str[9:]) // err

	i := strings.Index(str, "字符串") // or "字"
	fmt.Println(str[i:])

	runeStr := []rune(str)
	fmt.Println(string(runeStr[9:]))

	println()

	// 数字转换成字符串使用 strconv 处理，如果直接使用 string 则会被当作编码转义 char 处理
	fmt.Println(strconv.Itoa(100))
}

// 字符串拼接
func joint(s string) {
	fmt.Println(s + "拼接1")
	ss := fmt.Sprintf("%s%s", s, ",拼接2")
	fmt.Println(ss)
}
