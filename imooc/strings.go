package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱慕课网!！"             // utf8 占3个字节 字母占1个 符号占1个 中文符号占3个 // utf8
	fmt.Println(len(s))           // 22
	fmt.Printf("%s\n", []byte(s)) // 字符串转成切片  打印源内容
	//fmt.Printf("%X\n",[]byte(s)) // 字符串转成切片  打印源内容 %X 打出具体字节的数字
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch 就是 rune   // ch 变成了 Unicode 编码
		fmt.Printf("（%d %X）", i, ch)
	}
	fmt.Println()

	// 以utf8的编码方式数
	fmt.Println(utf8.RuneCountInString(s)) //10

	// 复杂方法
	bytes := []byte(s) //转切片
	fmt.Println("字符串转切片", bytes)
	for len(bytes) > 0 { // 如果切片长度大于0
		ch, size := utf8.DecodeRune(bytes) // 一个一个字符拿出来   1 字符  2 字符的size 中文3 英文 1
		fmt.Println("字符：", ch, "size：", size)
		bytes = bytes[size:]
		fmt.Printf("%c \n", ch) //%c 打印一个字符
	}
	fmt.Println()

	// 简单方法
	for i, ch := range []rune(s) { // rune 每个字符 占四个字节
		fmt.Printf("(%d %c)", i, ch)
	}

	// 使用[]byte获得字节
	// 使用len获得字节长度
	// 使用utf8.RuneCountInString 获取字符数量
	// 使用range 遍历pos,rune 对

	fmt.Println()
	// 其他字符串操作
	//Fields
	fmt.Println(strings.Fields("dsadas asdsa  asd  ads asd asd asd ")) // 输入一个字符串 按照空格 分割成一个切片
	//Split
	a := strings.Split("dsadas+asdsa+asd+ads+asd+asd+asd", "+")
	fmt.Println(a) // 按照指定的符号 分割字符串
	//Join
	fmt.Println(strings.Join(a, "+")) // 字符串按照指定符号拼接

	//Contains
	fmt.Println(strings.Contains("seafood", "foo")) //判断字符串 是否包含一个字符出串  反回真假值

	//Index
	fmt.Println(strings.Index("chicken", "ken")) // 返回substr 开始位置 不存在就返回-1
	//ToLower
	fmt.Println(strings.ToLower("Gopher")) // 将大写转小写

	//ToUpper
	fmt.Println(strings.ToUpper("Gopher")) // 将小写转大写

	//Trim
	fmt.Println(strings.Trim("¡¡¡Gopher!!!","¡!")) // 去两侧cutset
	//TrimRight
	fmt.Println(strings.TrimRight("¡¡!!Gopher¡¡!!","¡!")) // 去右侧cutset
	//TrimLeft
	fmt.Println(strings.TrimLeft("¡¡!!Gopher¡¡!!","¡!")) // 去左侧cutset
}
