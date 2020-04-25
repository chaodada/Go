package main

import (
	"fmt"
	"io/ioutil"
)

// 条件语句

func eval(a, b int, op string) int {
	var res int
	// switch
	switch op {
	case "+":
		res = a + b // 会自动加break
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	default:
		panic("unsupported operator:" + op) // 报错 让程序停下来
	}
	return res
}

// 不带表达式也可以 switch
func grade(score int) string {
	g := ""
	switch { // 不带表达式也可以 switch
	case score < 0 || score>100:
		panic(fmt.Sprintf("Wrong score: %d",score)) // 报错 让程序停下来
	case score < 60:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {

	// 写法1 if
	//const filename = "abc.txt" // 定义文件
	//contents,err:=ioutil.ReadFile(filename) // 读取文件 返回两个参数 第一个为内容 第二个为错误
	//if err!=nil {  // 如果错误存在
	//	fmt.Println(err) // 打印错误
	//}else {
	//	// 输出文件内容
	//	fmt.Printf("%s\n",contents)
	//}

	// 写法2 if
	// if 条件可以赋值 但是出了if 赋值变量就失效了
	const filename = "abc.txt"                                  // 定义文件
	if contents, err := ioutil.ReadFile(filename); err != nil { // 读取文件 返回两个参数 第一个为内容 第二个为错误
		fmt.Println(err) // 打印错误
	} else {
		// 输出文件内容
		fmt.Printf("%s\n", contents)
	}
	// fmt.Println(contents) 写法2 出了if 就不能用 contents 了

	// switch
	// switch 会自动加break 除非使用fallthrough
	fmt.Printf("%d\n", eval(1, 2, "+"))
	fmt.Printf("%d\n", eval(1, 2, "*"))

	fmt.Printf("%s\n",grade(6))
	//fmt.Println(grade(59),grade(101)) // 报错


}
