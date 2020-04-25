package main

import "fmt"

func main() {
	fmt.Println("Hello Word!")
}




// 1.1

// go build  编译
// go build -o 自定义的名字

// 在根目录 go install hello  就会在GOPATH 目录下的bin目录生成编译文件

// 1.2

// 所有的源码都是以.go结尾

//1.Go的函数、变量、常量、自定义类型、包(package)的命名方式遵循以下规则：
//1）首字符可以是任意的Unicode字符或者下划线
//2）剩余字符可以是Unicode字符、下划线、数字
//3）字符长度不限

