package main

import "fmt"

//定义一个 结构体
type TreeNode struct {
	value       int       // 内容
	left, right *TreeNode //类型是 指针
}

func (node TreeNode) print() { // (node TreeNode)  // 显示定义和命名方法的接收者
	fmt.Print(node.value, "    \n")
}

//func print(node TreeNode) { //  另一种写法
//	fmt.Print(node.value)
//}
// 设置值
func (node *TreeNode) setValue(value int) { // 传了一个指针   // 使用指针作为方法的接收者 只有指针才能改变结构内容  nil指针也可以调用方法
	if node == nil {
		fmt.Println("node to nil")
		return
	}
	node.value = value
}

// 遍历数
func (node *TreeNode) traverse() { // 中式便利     先左在中在右
	if node == nil {
		return
	}
	node.left.traverse() // 便利左侧树
	node.print()
	node.right.traverse() //便利右侧树

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

	root.print()
	root.setValue(100)
	root.print()

	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()
	root.print()

	// nil 也可以调用方法

	var psRoot *TreeNode
	fmt.Println(psRoot)
	psRoot.setValue(1)
	psRoot = &root
	psRoot.setValue(200)
	psRoot.print()

	fmt.Println()
	var roo TreeNode
	roo.value = 10
	roo.left = createTreeNode(20)
	roo.right = createTreeNode(30)
	roo.left.left = createTreeNode(40)
	roo.left.right = createTreeNode(50)
	roo.right.left = createTreeNode(60)
	roo.right.right = createTreeNode(70)

	roo.traverse()

	//值接收是拷贝

	//值接受者 vs 指针接收者
	//要改变内容必须使用指针接受者
	//结构过大也考虑使用指针接收者
	//一致性 有指针接收者 最好都是指针接收者

	//值接收者 是go 语言特有
	// 值/指针接收者 均可接收值 / 指针

	// 封装
	//名字一般
	//UserName
	//许多单词首字母大写连一块
	//首字母大写
	//public
	//首字母小写
	//private

}
