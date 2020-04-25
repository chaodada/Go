package main

import "fmt"

// 变量一旦定义就必须要用到

// 函数外边定义变量不能用 :=
// 作用域 包作用域
var aa = 3
var ss = "kkk"
var bb = true

// 可以直接括号包起来 不用写太多var
var (
	xx = 6
	dd = "dd"
)

// 定义变量
func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a,s)
	fmt.Printf("%d %q\n", a, s)
}

// 定义变量赋值
func variableInitialValue() {
	var a, b int = 3, 4 // 一次性定义多个变量 并且给初值
	var c, d int        // 一次性定义多个变量 不给初值
	var s string = "abc"
	fmt.Println(a, b, s)
	fmt.Println(c, d)
}

// 省略类型定义变量
func variableTypeDeducttion() {
	var a, b, c, d = 3, 4, true, "def"
	fmt.Println(a, b, c, d)
}

// 简化写法 :=
func variableShorter() {
	a, b, c, d := 3, 4, true, "def" // := 定义变量 赋值
	b = 6
	fmt.Println(a, b, c, d)
}

// 内建变量类型
// bool string
// 加u 无符号整数  不加 u 有符号整数  有符号整数分两类 1 规定长度 int8  int16  int32 int64   2 不规定长度(根据操作系统 32位就是32 64位就是64 )
// uintptr 指针
// byte 字节
// rune 字符型 32位

func main() {
	fmt.Println("Hello Word!")
	variableZeroValue()      // 调用函数
	variableInitialValue()   // 调用函数
	variableTypeDeducttion() // 调用函数
	variableShorter()        // 调用函数
	fmt.Println(aa, ss, bb, xx, dd)
}


