package main

import "fmt"

func main() {

	// 定义 map
	m := map[string]string{
		"name": "超大大",
		"site": "19year",
	}
	fmt.Println(m)

	// 定义map
	m2 := make(map[string]int) // 空的mao
	fmt.Println(m2)            // m2==empty map

	// 定义map
	var m3 map[string]int
	fmt.Println(m3) //m3 ==nil

	// 遍历map （跟数组差不多）
	for k, v := range m {
		fmt.Println("key", k, "val", v)
	}
	for _, v := range m {
		fmt.Println("val", v)
	}
	// 不能保证遍历顺序 如果需要顺序 需要手动对key排序

	//map 的key 是无序的 先后顺序会变得

	//map的操作
	//获取姓名
	name := m["name"]
	fmt.Println(name)
	// 获取不存在的key
	names := m["names"]
	fmt.Println(names) //值不存在的时候 返回map类型的初始值

	// 判断key在不在
	name, ok := m["name"]
	fmt.Println(name, ok)

	names, ok = m["names"]
	fmt.Println(names, ok)
	// 一般这样写
	if names, ok = m["names"]; ok {
		fmt.Println(names)
	} else {
		fmt.Println("key不存在")
	}

	// 删除元素
	fmt.Println(m)    //map[name:超大大 site:19year]
	delete(m, "name") //删掉name
	fmt.Println(m)    // map[site:19year]

	//使用len获取元素的个数
	size := len(m)
	fmt.Println(size)

	//map 的key
	//map 使用哈希表  必须可以比较相等
	//除了slice map function 的内建类型都可以作为key
	//Struct类型不包含上述字段 也可以作为key



}
