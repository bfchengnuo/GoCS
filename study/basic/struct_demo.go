package main

import "fmt"

// 定义结构体
type treeNode struct {
	value int
	// 避免复制，优先使用指针
	left, right *treeNode
}

// 组合的方式进行扩展
type myTreeNode struct {
	node *treeNode
}

// 定义结构体，测试匿名成员
type cusNode struct {
	// 因为没有名字，权限由具体类型控制（继承）
	treeNode
	name string
}

// 关于结构体的演示
func main() {
	// 定义
	var root treeNode
	// 对于结构体指针，可以直接使用『点』来调用
	// 类似是一个语法糖，会自动进行 (*t1).xxx 处理
	t1 := new(treeNode)
	t1.value = 1
	t2 := &treeNode{}
	t2.value = 2
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{4, nil, nil}
	// 指针可以直接用 . 调用
	root.right.left = new(treeNode)

	root.print()
	root.setVal(233)
	root.print()

	// 调用扩展的结构体
	extRoot := myTreeNode{&root}
	extRoot.echo()

	// 匿名成员测试
	cn := cusNode{root, "mps"} // 无法直接给 treeNode 中的属性赋值
	// 省去路径，直接调用，类似继承
	fmt.Println(cn.value)
	// 其他赋值写法
	cn = cusNode{treeNode{value: 1}, "mps"}
	cn = cusNode{
		treeNode: treeNode{1, nil, nil},
		name:     "mps",
	}
}

// 结构体增加方法
func (node treeNode) print() {
	fmt.Println(node.value)
}

// 使用指针接收者，使用上与值类型并无区别
// 1. 需要修改接收者中的值
// 2. 接收者是拷贝代价比较大的大对象
// 3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
func (node *treeNode) setVal(val int) {
	node.value = val
}

// 扩展方法
func (myNode *myTreeNode) echo() {
	if myNode == nil {
		return
	}
	left := myTreeNode{myNode.node.left}
	right := myTreeNode{myNode.node.right}

	fmt.Println("扩展...")
	left.node.print()
	myNode.node.print()
	right.node.print()
}

// 工厂函数创建，或者使用 new
// 可以当作构造方法用, 返回的指针可以正常使用
func createNode(val int) *treeNode {
	return &treeNode{value: val}
}

// 因为经常使用指针操作结构体，一般定义
func show() {
	pp := &treeNode{3, nil, nil}
	// 等价于
	// pp := new(treeNode)
	// *pp = treeNode{3, nil, nil}
	fmt.Println(pp)
}
