package main

import (
	"fmt"

	"golang.org/x/tools/container/intsets" //使用包
)

func TestSparse() {
	s := intsets.Sparse{}
	s.Insert(1)
	s.Insert(1000)
	s.Insert(1000000)
	fmt.Println(s.Has(1000))
	fmt.Println(s.Has(100000))
}

func main() {
	TestSparse()
}
