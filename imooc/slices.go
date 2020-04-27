package main

import "fmt"

func main() {
	//go语言中一般不直接使用数组
	//go中一般会使用切片
	//Slice 切片

	arr7 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr7[2:6] // 半开半闭区间 2 包含 6 舍弃 可以理解为php 数组分割array_slice()
	fmt.Println(s) // [2 3 4 5]


}
