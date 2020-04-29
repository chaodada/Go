package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	fmt.Println("值", s, "\n")

}

func main() {

	// 直接创建切片
	var s []int // 默认值为nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1) //每次切片装不下的时候cap 就会扩充1倍
	}
	fmt.Println(s)

	// 创建一个有值的切片
	s1 := []int{2, 4, 6, 8}
	printSlice(s1) // [2 4 6 8]

	// 创建一个已经知道长度的切片 但是不知道值
	s2 := make([]int, 16)     // len 长度是 16 cap 开辟一个 16 的数组
	printSlice(s2)            // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	s3 := make([]int, 10, 32) // len 长度是 10 cap 开辟一个 32 的数组
	printSlice(s3)            // [0 0 0 0 0 0 0 0 0 0]

	copy(s2, s1)   // 将s1 拷贝到 s2
	printSlice(s2) //  [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0]

	// 将s2 中的8 删除  （删除中间的元素）
	// s2[:3]+s2[4:] 不能这么写
	s2 = append(s2[:3], s2[4:]...) // 可变参数 ... 三个点就可以了
	printSlice(s2)                 //  [2 4 6 0 0 0 0 0 0 0 0 0 0 0 0]

	fmt.Println("1111111\n")

	// 删除头元素
	front := s2[0] // 2
	s2 = s2[1:]    //4 6 0 0 0 0 0 0 0 0 0 0 0 0
	fmt.Println("头元素", front, "\n")
	fmt.Println(s2, "\n")
	printSlice(s2)

	// 删除尾元素
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("尾元素", tail, "\n")
	fmt.Println(s2, "\n")
	printSlice(s2)

}
