package main

import "fmt"

func main() {


	// 寻找最长不含有重复字符的的字串
	// abcabcbb->abc
	// bbbbb->b
	// pwwkew->wke

	fmt.Println(lenthOfNonRepeatingSubStr("abcabcbb")) // 3
	fmt.Println(lenthOfNonRepeatingSubStr("bbbbb"))    // 1
	fmt.Println(lenthOfNonRepeatingSubStr("pwwkew"))   // 3
	fmt.Println(lenthOfNonRepeatingSubStr(""))         // 0
	fmt.Println(lenthOfNonRepeatingSubStr("b"))        // 1
	fmt.Println(lenthOfNonRepeatingSubStr("abcdef"))   // 6


	fmt.Println(lenthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(lenthOfNonRepeatingSubStr("一二三二一"))  

	//lenthOfNonRepeatingSubStr("ababc")
}

func lenthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int) // 创建一个map key 为byte类型 值为int类型
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) { // 将字符串转byte类型的切片   每一个ch就是字符串转的ascii码

		fmt.Println("字符串切片第", i, "个，内容为", ch)
		fmt.Println("当前的lastOccurred的内容为", lastOccurred)
		fmt.Println("当前的start的内容为", start)
		fmt.Println("当前的maxLength的内容为", maxLength)

		lastI, ok := lastOccurred[ch]
		fmt.Println("lastOccurred[",ch,"]的值", lastI, "是否存在", ok)

		if ok && lastI >= start { // lastOccurred 当在第一次循环的时候 是一个空的map 所以按照 字符串的第一个字符的ascii码取不出东西来 所以条件不成立
			fmt.Println("OK")
			start = lastOccurred[ch] + 1
		} else {
			fmt.Println("不OK")

		}
		if i-start+1 > maxLength { // 0 - 0 + 1 > 0
			maxLength = i - start + 1 // 0 - 0 + 1
		}
		lastOccurred[ch] = i // 0
		fmt.Println("当前的lastOccurred的内容为", lastOccurred)
		fmt.Println("当前的maxLength的内容为", maxLength)
		fmt.Println()
	}
	return maxLength
}
