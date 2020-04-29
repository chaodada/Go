package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := Queue{1}
	fmt.Println(q)
	q.Push(2)
	fmt.Println(q)
	q.Push(3)
	fmt.Println(q)
	t.Logf("test add succ")
}

func TestQueue_Pop(t *testing.T) {
	q := Queue{1, 2, 3}
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)
	q.Pop()
	t.Logf("test Sub succ")
}

func TestQueue_IsEmpty(t *testing.T) {
	q := Queue{1, 2, 3}
	fmt.Println(q.IsEmpty())
	q.Pop()
	q.Pop()
	q.Pop()
	fmt.Println(q.IsEmpty())
	t.Logf("test Sub succ")
}
