package main

import "fmt"

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

func main() {
	//int_value_2 := 8
	//fmt.Println(int8(int_value_2))

	//str := "Hello, 世界o"
	//for i, ch := range str {
	//	fmt.Println(i, ch) // ch 的类型为 rune
	//}

	//var c [3][3][3]float64
	//fmt.Println(c)
	//
	//var multi [9][9]string
	//for j := 0; j < 9; j++ {
	//	for i := 0; i < 9; i++ {
	//		n1 := i + 1
	//		n2 := j + 1
	//		if n1 < n2 { // 摒除重复的记录
	//			continue
	//		}
	//		multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
	//	}
	//}
	//
	//fmt.Println(multi)
	//
	//// 打印九九乘法表
	//for _, v1 := range multi {
	//	for _, v2 := range v1 {
	//		fmt.Printf("%-8s", v2) // 位宽为8，左对齐
	//	}
	//	fmt.Println()
	//}

	// 先定义一个数组
	//months := [...]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	//
	//// 基于数组创建数组切片
	//q2 := months[3:6]     // 第二季度
	//summer := months[5:8] // 夏季
	//
	//all := months[:]         // 基于数组 全创建成切片
	//firsthalf := months[:6]  // 上半年
	//secondhalf := months[6:] // 下半年
	//
	//fmt.Println(q2)
	//fmt.Println(summer)
	//fmt.Println(all)
	//fmt.Println(firsthalf)
	//fmt.Println(secondhalf)
	//fmt.Println(len(q2)) // 切片测长度
	//fmt.Println(cap(q2)) // 切片的容量
	//
	//q1 := firsthalf[1:3] // 基于firsthalf的前3个元素构建新数组切片
	//
	//fmt.Println(q1)
	//fmt.Println(len(q1)) // 切片测长度
	//fmt.Println(cap(q1)) // 切片的容量
	//
	//slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//slice3 = slice3[:len(slice3)-5] // 删除 slice3 尾部5个元素
	//fmt.Println(slice3)
	//slice3 = slice3[5:] // 删除 slice3 头部 5 个元素
	//fmt.Println(slice3)

	//score := 100
	//switch {
	//case score >= 90:
	//	fmt.Println("Grade: A")
	//case score >= 80 && score < 90:
	//	fmt.Println("Grade: B")
	//case score >= 70 && score < 80:
	//	fmt.Println("Grade: C")
	//case score >= 60 && score < 70:
	//	fmt.Println("Grade: D")
	//default:
	//	fmt.Println("Grade: F")
	//}

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}
JLoop:
	fmt.Println("啊啊啊")
}
