package main

import "fmt"

// 数组
func main() {
	var arr1 [5]int // 定义数组 数组中存在5个数字
	arr2 :=[3]int{1,3,5} // := 必须写初值
	arr3 :=[...]int{1,2,3,4,5,6,7} // ... 让编译器 来数数组中有多少个

	fmt.Println(arr1,arr2,arr3)

	var grid [4][5]int // 定义一个二维数组 大数组包含4个小数组 每个小数组包含5个内容
	fmt.Println(grid)

	var grids [2][3][4] int
	fmt.Println(grids)






}
