package main

import (
	"fmt"
	"io/ioutil"
)

// 条件语句

func main() {
	const filename = "abc.txt" // 定义文件
	contents,err:=ioutil.ReadFile(filename) // 读取文件 返回两个参数 第一个为内容 第二个为错误
	if err!=nil {  // 如果错误存在
		fmt.Println(err) // 打印错误
	}else {
		// 输出文件内容
		fmt.Printf("%s\n",contents)
	}

}
