package main

import (
	"fmt"
	"imooc/jiekou/mock"
	"imooc/jiekou/real"
	"time"
)

// 接口

// 定义接口
type Retiever interface {
	Get(url string) string
	Print()
}

func download(r Retiever) string {
	return r.Get("http://www.baidu.com")
}

func main() {

	// Go 接口
	// duck typing 像鸭子就是鸭子 (描述事物外部行为而非内部结构)
	// 大黄鸭是鸭子吗 有生命 是鸭子生的。。。

	// go 属于结构化类型系统 类似duck typing

	// python
	//def download(res):
	//	return res.get("baidu.com")
	//运行时才知道 res 有没有get
	//需要注释来说明 接口

	//c++
	//template <class R>
	//string atom.Download( const R & res){
	//	return res.get("baidu.com")
	//}
	//编译时才知道传入的res 有没有get
	//需要注释来说明 接口

	//java
	//<R extends Res>
	//	String downlod(R r){
	//	return r.get("baidu.com")
	//}
	// 传入参数必须实现 Res接口
	// 不是 duck typing

	//go 语言的 duck typing
	//如果 同时需要 实现两个接口怎么办
	//同时 具有python c++ 的duck typing 的灵活性
	//又具有java的类型检查

	// 接口的定义  实现

	//download //使用者
	//res      //实现者

	//go的接口由使用者定义

	var r Retiever
	r = mock.Retriever{"这是内容"}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	//r.Print()

	r = &real.URetriever{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)
	//fmt.Println(download(r))
	//r.Print()
	inspect(r)
}

func inspect(r Retiever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("sssss", v.Contents)
	case *real.URetriever:
		fmt.Println("bbbbb", v.UserAgent)
	}
}
