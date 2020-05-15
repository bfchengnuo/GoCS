package main

import (
	"fmt"
	"go/types"
)

// 参考 gin 框架
type IRouter interface {
	test()
}

type RouterGroup struct{}

// impl
func (x RouterGroup) test() {
	fmt.Println("Impl")
}

func main() {
	// 让编译器检查类型是否实现接口
	var _ IRouter = &RouterGroup{}

	// 使用类型断言检查
	var x interface{}
	x = "Hello mps"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}

// 使用 switch 判断类型
func swCheck(t interface{}) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}

	// 另一种写法
	switch x := t.(type) {
	case types.Nil:
		fmt.Println("Nil")
	case int:
		fmt.Println("int - ", x)
	case string:
		fmt.Println("string - ", x)
	default:
		fmt.Println("default")
	}
}
