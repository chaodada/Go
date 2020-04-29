package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//函数 没有默认参数 可选参数

// 函数
func calculator(a, b int, op string) (int, error) { // 这个函数只有2 个返回值 int 类型 error 类型
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		//return a / b
		a, _ := div(a, b) //  但是在/ 的时候 调用了 div 函数 有两个返回值 可以利用 _ 忽略掉不用的返回值
		return a, nil
	default:
		//panic("报错：" + op)
		return 0, fmt.Errorf("报错：%s 运算符号不正确", op)
	}
}

// 函数返回多个值
// 除法 13 /3 =4 。。。。。。1

// 多个返回值的用法 一般用法 通常 返回一个值 一个err

func div(a, b int) (c, d int) { // (int, int) 可以只声明多个变量的类型  也可以声明名字 (c ,d int)
	//return a / b, a % b   另一种写法
	c = a / b
	d = a % b
	// return c,d 另一种写法
	return
}

// 函数式编程  函数的参数也可以是一个函数
func apply(op func(int, int) int, a, b int) int { // 第一个参数 一个函数(该函数接受两个int数字 返回一个int 数字) 第二三个参数 int 类型的数字  返回值是 int

	p := reflect.ValueOf(op).Pointer()    // 映射
	opName := runtime.FuncForPC(p).Name() // 获取函数名

	fmt.Printf("函数名 %s "+
		"参数为(%d,%d)", opName, a, b)
	return op(a, b) // 相当于执行的
}

// math.Pow 函数只能接收  float64  所以重写 pow 函数 转成int
// 计算a 的 b 次方
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
//所有参数累加
func sum(numbers ...int) int {
	s := 0
	for i := range numbers { // i 等于 参数中的每一个数
		s += numbers[i] // 所有参数累加
	}
	return s
}

// 测试指针a
func zhizhena(n *int) {

	*n += 1
	//return *n
}

func main() {
	// 计算器
	fmt.Println(calculator(1, 2, "+"))  // 0  nil
	fmt.Println(calculator(1, 2, "*"))  // 0  nil
	fmt.Println(calculator(10, 3, "/")) // 0  nil
	fmt.Println(calculator(10, 3, "x")) // 返回0 跟error
	aa, err := calculator(10, 3, "x")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aa)
	}
	// 多个返回值函数
	b, d := div(13, 2)
	fmt.Println(b, d)

	//fmt.Println(apply(pow,3,4)) // 计算3 的4 次方
	// 另一种写法 匿名函数
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	// 所有参数累加
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))

	// go 指针
	// 指针不能运算
	var a int = 1
	var pa *int = &a // pa 相当于a 的引用
	*pa += 2         // 相当于 a=2
	fmt.Println(a)

	// 参数传递
	// 值传递 ？ （可以理解为将值传到函数中 函数外不会发生改变）  引用传递？（可以理解为将值的地址传到函数中 函数外会根据函数处理逻辑而发生改变）
	// go 语言只有值传递 一种方式

	// 值传递 与指针配合
	var zz int = 10
	fmt.Println(zz) // 10
	zhizhena(&zz)   // 函数 里边加1
	fmt.Println(zz) // 11  函数影响到外边的变量 (因为传递的是引用 如果是值传递则不会发生影响)

	a, b = 100, 200
	fmt.Println(a, b) //100, 200
	swapa(&a, &b)
	fmt.Println(a, b) //200, 100
	a, b = swapb(a, b)
	fmt.Println(a, b) //100, 200

}

// 数据调换位置a
func swapa(a, b *int) { // 接受两个指针类型的参数 &a &b
	fmt.Println(a)
	fmt.Println(b)
	*b, *a = *a, *b // *b=*a *a=*b
}

// 数据调换位置 b
func swapb(a, b int) (int, int) { // 接受两个int类型的参数 a b
	return b, a // 这样也可实现 调换位置
}
