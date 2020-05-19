package test

import (
	"reflect"
	"testing"
)

// 测试函数，必须 Test 开头, 必须接收一个 *testing.T 类型参数
// 可以在当前包路径下执行 go test [-v] [-cover]
func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"} // 期望的结果
	// 因为slice不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}

	// 子测试
	//t.Run()
}
