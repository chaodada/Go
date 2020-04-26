package main

import "fmt"

// 函数
func calculator(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("报错：" + op)
	}
}

func main() {
	fmt.Println(calculator(1, 2, "+"), calculator(1, 2, "/"))
}
