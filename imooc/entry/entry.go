package main

import (
	"fmt"
	"tree"
)

func main() {

	// go 语言面向对象  go 只有封装 没有继承--多态
	// go 没有class 只有 struct

	var root tree.TreeNode
	//root2:=treeNode{}

	fmt.Println(root) // {0 <nil> <nil>}

	root = tree.TreeNode{Value: 3}

	fmt.Println(root) // {3 <nil> <nil>}

	root.Left = &tree.TreeNode{}
	fmt.Println(root) // {3 0xc0000a6060 <nil>}

	root.Right = &tree.TreeNode{5, nil, nil}
	fmt.Println(root) // {3 0xc0000a6060 0xc0000a60a0}

	root.Right.Left = new(tree.TreeNode)           //  不论地址还是结构体本身 一律使用.来访问成员
	root.Right.Right = &tree.TreeNode{6, nil, nil} //
	fmt.Println(root)                              //

	// 结构没有构造方法  有工厂函数
	root.Left.Right = tree.CreateTreeNode(2) // 使用自定义工厂函数
	root.Left.Left = tree.CreateTreeNode(2)

	fmt.Println(root) //{3 0xc00000c0c0 0xc00000c100}

	//      3
	//   /    \
	//  0      5
	// / \    / \
	//2   2   0  6
	//以上代码 就是 创建一个树

	// 为结构定义方法
	fmt.Println()
	root.Print() // 调用结构体的自定义方法
	fmt.Println()
	//print(root) // 另一种方法
	fmt.Println()

	root.Right.Right.Print()
	fmt.Println()
	root.Right.Right.SetValue(8)
	fmt.Println()
	root.Right.Right.Print()
	fmt.Println()

	root.Print()
	root.SetValue(100)
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()
	root.Print()

	// nil 也可以调用方法

	var psRoot *tree.TreeNode
	fmt.Println(psRoot)
	psRoot.SetValue(1)
	psRoot = &root
	psRoot.SetValue(200)
	psRoot.Print()

	fmt.Println()
	var roo tree.TreeNode
	roo.Value = 10
	roo.Left = tree.CreateTreeNode(20)
	roo.Right = tree.CreateTreeNode(30)
	roo.Left.Left = tree.CreateTreeNode(40)
	roo.Left.Right = tree.CreateTreeNode(50)
	roo.Right.Left = tree.CreateTreeNode(60)
	roo.Right.Right = tree.CreateTreeNode(70)

	roo.Traverse()

	//值接收是拷贝

	//值接受者 vs 指针接收者
	//要改变内容必须使用指针接受者
	//结构过大也考虑使用指针接收者
	//一致性 有指针接收者 最好都是指针接收者

	//值接收者 是go 语言特有
	// 值/指针接收者 均可接收值 / 指针

}
