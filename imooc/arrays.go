package main

import "fmt"

// 数组
// range 关键字 可以返回 k v
func main() {
	var arr1 [5]int                       // 定义数组 数组中存在5个数字
	arr2 := [3]int{1, 3, 5}               // := 必须写初值
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7} // ... 让编译器 来数数组中有多少个

	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int // 定义一个二维数组 大数组包含4个小数组 每个小数组包含5个内容
	fmt.Println(grid)

	var grids [2][3][4]bool // 定义一个多维数组 大数组包含2个小数组 每个小数组包含3个数组 每个数组包含4个元素
	fmt.Println(grids)

	// 遍历数组
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 方法一
	for i := 0; i < len(arr4); i++ {
		fmt.Printf("%d\n", arr4[i])
	}
	// 方法二
	for i := range arr4 {
		fmt.Printf("%d\n", arr4[i])
	}
	fmt.Printf("分割行\n")
	// 遍历多维数组
	var gridss [2][3][4]bool    // 定义一个多维数组 大数组包含2个小数组 每个小数组包含3个数组 每个数组包含4个元素
	for i, vi := range gridss { // i 下标 v 值
		fmt.Println(i, vi)
		for j, vj := range gridss[i] {
			fmt.Println(j, vj)
			for x, xj := range gridss[i][j] {
				if x%2 == 0 {
					xj = true
				}
				fmt.Println(j, xj)
			}
		}
	}
	fmt.Println(gridss)

	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for key, value := range arr5 { // 忽略k 只要v 将key改写成_
		fmt.Println(key, value)
	}

	fmt.Printf("分割行\n")
	// 如果数组中的值大于maxValue 就拿到他的key 以及value
	numbers := [3]int{1, 2, 3}
	maxValue := 2
	for i, v := range numbers {
		if v > maxValue {
			fmt.Printf("v---%d大于maxValue---%d，他的下标为---%d\n", v, maxValue, i)
		}
	}

	// 为什么用range 遍历
	// 1 意图明确


	// 数组是值类型

}
