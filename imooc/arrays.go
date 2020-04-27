package main

import (
	"fmt"
)

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
					xj = true // 外边不生效
					gridss[i][j][x] = true
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

	// 数组是值类型 遍历数组的时候修改了数组内容 在遍历外边打印数组 还是原来的内容 不会改变 // 就类似与php foreach 遍历的时候没有加&
	printArray([5]int{1, 2, 3, 4, 5}) // 正确
	// printArray([6]int{1, 2, 3, 4, 5}) // 报错 // 原因 [6]int 与 [5]int Go认为不是一个类型

	//如果想在数组遍历的时候修改数组的内容
	// 可以使用指针
	arr6 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr6) // [1 2 3 4 5]
	editArray(&arr6)  // 传递指针
	fmt.Println(arr6) //[2 3 4 5 6]
	for key, _ := range arr6 {
		arr6[key] += 1 //此处注意 如果是直接用遍历出来的值赋值  在遍历外部 原数组不会发生改变   如果使用key 的方式修改内容 则愿数组会发生改变
	}
	fmt.Println(arr6) //[3 4 5 6 7]
	for _, value := range arr6 {
		value += 1 //此处注意 如果是直接用遍历出来的值赋值  在遍历外部 原数组不会发生改变   如果使用key 的方式修改内容 则愿数组会发生改变
	}
	fmt.Println(arr6) //[3 4 5 6 7]


}

func editArray(arr *[5]int) {
	for key, _ := range arr {
		arr[key] += 1
	}
}

func printArray(arr [5]int) { //这个函数只能接收长度为5的数组 其他长度都会报错
	for key, value := range arr {
		fmt.Println(key, value)
	}
}
