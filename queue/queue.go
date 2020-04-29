package queue

import "fmt"

type Queue []int // 切片

// 入队列
func (q *Queue) Push(v int) {
	fmt.Println("1123", q)
	*q = append(*q, v)
}

// 出队列
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
