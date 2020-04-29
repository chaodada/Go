package tree

import "fmt"

//定义一个 结构体
type TreeNode struct {
	Value       int       // 内容
	Left, Right *TreeNode //类型是 指针
}

func (node TreeNode) Print() { // (node TreeNode)  // 显示定义和命名方法的接收者
	fmt.Print(node.Value, "    \n")
}

//func print(node TreeNode) { //  另一种写法
//	fmt.Print(node.value)
//}
// 设置值
func (node *TreeNode) SetValue(value int) { // 传了一个指针   // 使用指针作为方法的接收者 只有指针才能改变结构内容  nil指针也可以调用方法
	if node == nil {
		fmt.Println("node to nil")
		return
	}
	node.Value = value
}

// 遍历数
func (node *TreeNode) Traverse() { // 中式便利     先左在中在右
	if node == nil {
		return
	}
	node.Left.Traverse() // 便利左侧树
	node.Print()
	node.Right.Traverse() //便利右侧树

}

// 有工厂函数
func CreateTreeNode(value int) *TreeNode {
	return &TreeNode{Value: value} // 注意返回了局部变量的地址
}

// 封装
//名字一般
//UserName
//许多单词首字母大写连一块
//首字母大写
//public
//首字母小写
//private

// 为结构定义的方法必须放在同一个包内
// 可以是不同的文件
//如何扩充系统类型 或者别人的类型
