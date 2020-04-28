package main

import "fmt"

func main() {
	//go语言中一般不直接使用数组
	//go中一般会使用切片
	//Slice 切片
	//切片本身是没有数据的 就是对底层数组的一个视图   可以理解为创建一个切片就是 对数组的一部分数据（或者整个数组）进行引用

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 都是切片
	s := arr[2:6]                         // 半开半闭区间 2 包含 6 舍弃(也就是说6前一个) 可以理解为php 数组分割array_slice() //下标从0开始取，，，  区间取左不取右
	fmt.Println(s)                        // [2 3 4 5]
	fmt.Println("arr[:6]", arr[:6], "\n") //  [0 1 2 3 4 5]
	fmt.Println("arr[2:]", arr[2:], "\n") // [2 3 4 5 6 7 8 9]
	fmt.Println("arr[:]", arr[:], "\n")   // [0 1 2 3 4 5 6 7 8 9]

	// Slice 是一个视图
	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("执行updateSlice(s1)")
	fmt.Println("s1 执行函数前=", s1) // [2 3 4 5 6 7 8 9]
	updateSlice(s1)
	fmt.Println("s1 执行函数后=", s1) //[100 3 4 5 6 7 8 9]
	fmt.Println("创建切片的原数组", arr) //[0 1 100 3 4 5 6 7 8 9]

	fmt.Println("执行updateSlice(s2)")
	fmt.Println("s2 执行函数前=", s2) //  [0 1 100 3 4 5 6 7 8 9]
	updateSlice(s2)
	fmt.Println("s2 执行函数后=", s2) // [100 1 100 3 4 5 6 7 8 9]
	fmt.Println("创建切片的原数组", arr) // [100 1 100 3 4 5 6 7 8 9]

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array 执行函数前=", array) //[1 2 3 4 5 6 7 8 9]
	editArrayb(array[:])               // 传入数组的切片
	fmt.Println("array 执行函数后=", array) // [2 3 4 5 6 7 8 9 10]

	array2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("创建切片的原数组", array2) // [1 2 3 4 5 6 7 8 9]
	ar2 := array2[:]
	fmt.Println("创建切片[:]", ar2) // [1 2 3 4 5 6 7 8 9]
	ar2 = ar2[:5]
	fmt.Println("创建切片[:5]", ar2) // [1 2 3 4 5]
	ar2 = ar2[2:]
	fmt.Println("创建切片[2:]", ar2) // [3 4 5]

	// 坑 慢慢理解 (切片的扩展)
	fmt.Println("\n")
	array3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = array3[2:6]
	fmt.Println("创建切片[2:6]", s1) // [2 3 4 5]
	s2 = s1[3:5]
	fmt.Println("创建切片[3:5]", s2) //[5 6]
	s3 := s2[1:3]
	fmt.Println("创建切片[1:3]", s3) //[[6 7]

	// s3 = s2[1:3]                     0 1 <---s3的下标注意1(实际1对应array3的 7)
	// s2 = s1[3:5]                   0 1 *  <---s2的下标注意2(实际2对应array3的 7)
	// s1 = array3[2:6]         0 1 2 3 * * <---s1的下标注意4 5(实际4 5对应array3的 6 7 )
	// array3               0 1 2 3 4 5 6 7 <---数组中的数值

	//            [6 7] s3
	//		    [5 6 *] s2
	//	  [2 3 4 5 * *] s1
	//[0 1 2 3 4 5 6 7] array3

	// 切片的实现 ptr(切片开头的第一个元素) len(切片的长度 使用切片取值的时候如果大于等于len 就会报错 下标越界) cap(整个数组从ptr到结束的长度 切片扩展的时候就是这个特性 只要不超过cpa)

	// cap 从第一个|#| 到最后一个 |+| 的长度
	// 第一个|#| 称为ptr 切片s1 有5个元素 len5
	// s1=array[3:8]  即|#|#|#|#|#|
	// array  |*|*|*|#|#|#|#|#|+|+|+|

	// 切片可以向后扩展但是不可以向前扩展
	// 比如
	// s1=[2 3 4 5] s2 再怎么向前扩展 也只能是从2 开始
	// s[i] i不能超过len(s)（长度）    向后扩展不能超过底层数组cap(s)（长度）

	fmt.Println("\n")
	array3 = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("创建数组array3", array3) //

	s1 = array3[2:6]
	fmt.Println("创建切片s1[2:6]", s1) // [2 3 4 5]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 = s1[3:5]
	fmt.Println("创建切片s2[3:5]", s2) //[5 6]
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 = s2[1:3]
	fmt.Println("创建切片s3[1:3]", s3) //[6 7]
	fmt.Printf("s3=%v,len(s3)=%d,cap(s3)=%d\n", s3, len(s3), cap(s3))

	fmt.Println("\n")

	//向切片添加元素
	array4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("array4=", array4)

	s1 = array4[2:6]
	fmt.Println("s1=array4[2:6]", s1) // 2 3 4 5
	fmt.Println("array4=", array4)    // [0 1 2 3 4 5 6 7]

	s2 = s1[3:5]
	fmt.Println("s2 = s1[3:5]", s2) // 5 6
	fmt.Println("array4的值", array4)

	s3 = append(s2, 10)
	fmt.Println(s3)                     // 5 6 10
	fmt.Println(s2)                     // 5 6
	fmt.Println(s1)                     // 2 3 4 5
	fmt.Printf("array4的值为%v\n", array4) // [0 1 2 3 4 5 6 10]

	s4 := append(s3, 11)                // 已经超过array的cap的长度  系统会重新分配更大的底层数组
	fmt.Println(s4)                     // 5 6 10 11
	fmt.Println(s3)                     // 5 6 10
	fmt.Println(s2)                     // 5 6
	fmt.Println(s1)                     // 2 3 4 5
	fmt.Printf("array4的值为%v\n", array4) // [0 1 2 3 4 5 6 10]

	s5 := append(s4, 12) // 已经超过array的cap的长度  系统会重新分配更大的底层数组
	fmt.Println(s5)      // 5 6 10 11 12
	fmt.Println(s4)      // 5 6 10 11
	fmt.Println(s3)      // 5 6 10
	fmt.Println(s2)      // 5 6
	fmt.Println(s1)      // 2 3 4 5

	fmt.Printf("array4的值为%v\n", array4) // [0 1 2 3 4 5 6 10]

	//s4 s5 已经不是array4的视图

	// 添加元素时如果超越cap 系统会重新分配更大的底层数组
	// 由于值传递的关系 必须接收append的返回值
	// append() 会导致 切片 ptr  len  cap 发生改变
	// s=append(切片,值)





}

func editArrayb(arr []int) { // 传入一个切片
	for key, _ := range arr {
		arr[key] += 1
	}
}

// 参数传递一个切片
func updateSlice(s []int) {
	s[0] = 100
}
