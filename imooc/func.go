package main

import "fmt"

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

func main() {
	// 计算器
	fmt.Println(calculator(1, 2, "+"))  // 0  nil
	fmt.Println(calculator(1, 2, "*"))  // 0  nil
	fmt.Println(calculator(10, 3, "/")) // 0  nil
	fmt.Println(calculator(10, 3, "x")) // 返回0 跟error

	// 多个返回值函数
	b, d := div(13, 2)
	fmt.Println(b, d)
}
