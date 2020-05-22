package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 从标准输入生成读对象
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	text, _ := reader.ReadString('\n') // 读到换行

	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
