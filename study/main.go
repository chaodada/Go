package main

//一个结构体（struct）就是一组字段（field）。
type Vertex struct {
	int1 int
	int2 int
}

//结构体文法通过直接列出字段的值来新分配一个结构体。
//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
//特殊的前缀 & 返回一个指向结构体的指针。
var (
	v1 = Vertex{1, 2}    // 创建一个 Vertex 类型的结构体
	v2 = Vertex{int1: 1} // int2:0 被隐式地赋予
	v3 = Vertex{}        // int1:0 int2:0
	p  = &Vertex{1, 2}   // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {

	//Go 拥有指针。指针保存了值的内存地址。
	//
	//类型 *T 是指向 T 类型值的指针。其零值为 nil。
	//
	//var p *int
	//& 操作符会生成一个指向其操作数的指针。
	//
	//i := 42
	//p = &i
	//* 操作符表示指针指向的底层值。
	//
	//fmt.Println(*p) // 通过指针 p 读取 i
	//*p = 21         // 通过指针 p 设置 i
	//这也就是通常所说的“间接引用”或“重定向”。
	//
	//与 C 不同，Go 没有指针运算。

	//i, j := 42, 2701
	//
	//p := &i         // 指向 i
	//fmt.Println(*p) // 通过指针读取 i 的值  42
	//*p = 21         // 通过指针设置 i 的值
	//fmt.Println(i)  // 查看 i 的值 21
	//
	//p = &j         // 指向 j
	//*p = *p / 37   // 通过指针对 j 进行除法运算
	//fmt.Println(j) // 查看 j 的值 73

	//一个结构体（struct）就是一组字段（field）。
	//fmt.Println(Vertex{1,2})

	//结构体字段使用点号来访问。
	//v := Vertex{1, 2}
	//v.int1 = 4
	//fmt.Println(v.int1)

	//结构体字段可以通过结构体指针来访问。
	//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	//v := Vertex{1, 2}
	//p := &v
	//p.int2= 55
	//fmt.Println(v)

	//结构体文法通过直接列出字段的值来新分配一个结构体。
	//使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）
	//特殊的前缀 & 返回一个指向结构体的指针。
	//fmt.Println(v1)
	//fmt.Println(v2)
	//fmt.Println(v3)
	//fmt.Println(p)
	//p.int1=22
	//fmt.Println(p)

	//数组
	//类型 [n]T 表示拥有 n 个 T 类型的值的数组。
	//表达式
	//var a [10]int
	//会将变量 a 声明为拥有 10 个整数的数组。
	//数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
	//var a[11] int
	//a[0] = 1
	//a[1] = 2
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)
	//

	//切片
	//每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。
	//类型 []T 表示一个元素类型为 T 的切片。
	//切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
	//a[low : high]
	//它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
	//以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：
	//a[1:4]

	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//fmt.Println(primes)
	//var a []int = primes[0:4]
	//fmt.Println(a)




	//切片就像数组的引用
	//切片并不存储任何数据，它只是描述了底层数组中的一段。
	//更改切片的元素会修改其底层数组中对应的元素。
	//与它共享底层数组的切片都会观测到这些修改。

	//names := [4]string{
	//	"AAA",
	//	"BBB",
	//	"CCC",
	//	"DDD",
	//}
	//fmt.Println(names)
	//// 创建一个切片
	//a := names[0:2] // AAA  BBB
	//b := names[1:3] // BBB  CCC
	//fmt.Println(a)
	//fmt.Println(b)
	//b[0] = "XXX"
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(names)











}
