package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

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
// float32 float64   浮点数类型   complex64 complex128 (复数类型)

// 不懂
func euler() {
	//c:=3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
// 类型转换是强制的
func triangle() {
	// 勾股定理
	var a, b int = 3, 4
	var c int
	// 强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}

// 常量定义  也可以定义在函数外边 也就是包内部
const filenamea string = "111.txt" // 定义常量 也可以规定类型
// 也可以这样定义
const (
	aaa int = 66
	bbb int = 88
)

func consts() {
	const filename string = "abc.txt" // 定义常量 也可以规定类型
	const a, b = 3, 4                 // 定义常量 不规定类型 （常量数值可以做为各种类型使用）
	var c int
	c = int(math.Sqrt(a*a + b*b)) //  作为各种类型使用
	fmt.Println(filename, c)      // 函数内部的常量
	fmt.Println(filenamea)        // 函数外边的常量
	fmt.Println(aaa, bbb)         // 函数外边的常量

}

// 特殊的常量 枚举类型
func enums() {
	// 1
	const (
		php    = 0
		java   = 1
		golang = 2
	)
	fmt.Println(php, java, golang)
	// 2
	const (
		phpa    = iota // 自增值 0
		_              // 跳过1
		javaa          // 2
		golanga        // 3...
	)
	fmt.Println(phpa, javaa, golanga) // 0 2 3
	// 3
	const (
		b  = 1 << (10 * iota) // 1 * (10 * 0) // // 1 << 0   = 1 任何非0数的0次方都是1
		kb                    // 1 * (10 * 1) // // 1 << 10  = 1*2 *2 *2 *2 *2 *2 *2 *2 *2 *2
		mb                    // 1 * (10 * 2) // // 1 << 20  = 1*2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2
		gb                    // 1 * (10 * 3) // // 1 << 30  = 1*2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2
		tb                    // 1 * (10 * 4) // // 1 << 40  = 1*2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2
		pb                    // 1 * (10 * 5) // // 1 << 50  = 1*2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2*2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2 *2
	)
	fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824 1099511627776 1125899906842624
}

// 变量类型 写在变量名之后
// 没有char 类型 只有rune
// 原生支持复数类型

func main() {
	fmt.Println("Hello Word!")
	variableZeroValue()      // 调用函数
	variableInitialValue()   // 调用函数
	variableTypeDeducttion() // 调用函数
	variableShorter()        // 调用函数
	fmt.Println(aa, ss, bb, xx, dd)

	// 不懂  验证euler
	euler()

	// 勾股定理
	triangle()
	// 常量定义
	consts()
	// 特殊的常量 枚举类型
	enums()
}
