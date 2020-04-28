package main

import "fmt"

//定义一个 结构
type treeNode struct {
	value       int       // 内容
	left, right *treeNode //类型是 指针
}

func main() {

	// go 语言面向对象  go 只有封装 没有继承 多态

	var root treeNode
	fmt.Println(root) // {0 <nil> <nil>}
	root =treeNode{value: 3}
	fmt.Println(root) // {0 <nil> <nil>}
	root2.left=&treeNode{}
	root2.right=&treeNode{5,nil,nil}




}
