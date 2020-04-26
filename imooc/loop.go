package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// go 的循环

// 整数转二进制表达式
func convertToBin(n int) string {
	// 13对2取模 余1 商6
	// 6 对2取模 余0 商3
	// 3 对2取模 余1 商1
	// 1 对2取模 余1 商0
	result := ""
	//　循环省略初始条件
	for ; n > 0; n /= 2 { // n /= 2  等于  n=n/2
		//fmt.Printf("%d\n", n)
		lsb := n % 2                        // 取模
		result = strconv.Itoa(lsb) + result // strconv.Itoa()  int转字符串
	}
	return result
}

// 读取文件
func printFile(filename string) {
	file, err := os.Open(filename) // 打开文件
	if err != nil {
		panic(err) //终止程序 打印报错
	}
	scanner := bufio.NewScanner(file)
	// 循环省略初始条件 省略递增条件 这时候可以省略分号  相当于while循环
	for ; scanner.Scan(); {
		fmt.Println(scanner.Text()) // 读取文件
	}
}
// 死循环
func forever()  {
	for  { // for 条件全部省略  就是死循环
		fmt.Println("abc\n")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  //二进制为 101
		convertToBin(13), //二进制为 1101
	)
	printFile("abc.txt")
	forever();
}
