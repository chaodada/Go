package main

import "fmt"

//定义一个 结构体
type TreeNode struct {
	value       int       // 内容
	left, right *TreeNode //类型是 指针
}

func (node TreeNode) print() { // (node TreeNode)
	fmt.Print(node.value)
}

//func print(node TreeNode) { //  另一种写法
//	fmt.Print(node.value)
//}
// 设置值
func (node *TreeNode) setValue(value int) {
	node.value = value
}

// 有工厂函数
func createTreeNode(value int) *TreeNode {
	return &TreeNode{value: value} // 注意返回了局部变量的地址
}

func main() {

	// go 语言面向对象  go 只有封装 没有继承--多态
	// go 没有class 只有 struct

	var root TreeNode
	//root2:=treeNode{}

	fmt.Println(root) // {0 <nil> <nil>}

	root = TreeNode{value: 3}

	fmt.Println(root) // {3 <nil> <nil>}

	root.left = &TreeNode{}
	fmt.Println(root) // {3 0xc0000a6060 <nil>}

	root.right = &TreeNode{5, nil, nil}
	fmt.Println(root) // {3 0xc0000a6060 0xc0000a60a0}

	root.right.left = new(TreeNode)           //  不论地址还是结构体本身 一律使用.来访问成员
	root.right.right = &TreeNode{6, nil, nil} //
	fmt.Println(root)                         //

	// 结构没有构造方法  有工厂函数
	root.left.right = createTreeNode(2) // 使用自定义工厂函数
	root.left.left = createTreeNode(2)

	fmt.Println(root) //{3 0xc00000c0c0 0xc00000c100}

	//      3
	//   /    \
	//  0      5
	// / \    / \
	//2   2   0  6
	//以上代码 就是 创建一个树

	// 为结构定义方法
	fmt.Println()
	root.print() // 调用结构体的自定义方法
	fmt.Println()
	//print(root) // 另一种方法
	fmt.Println()



	root.right.right.print()
	fmt.Println()
    root.right.right.setValue(8)
	fmt.Println()
	root.right.right.print()
	fmt.Println()
	//root2.right=&treeNode{5,nil,nil}

	//fmt.Println()
	//nodes := []TreeNode{
	//	{value: 5},
	//	{},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	//fmt.Println()
	//apps:=TreeNode{}
	////app.value=5
	//app.left.right=createTreeNode(8)
	//fmt.Println(app)

}
