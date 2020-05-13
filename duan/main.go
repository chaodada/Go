package main

import "duan/web" // 引入web应用包

func main() {

	// 实例化应用
	app := &web.App{}

	// 初始化应用
	app.Initalize()

	// 运行应用
	app.Run("127.0.0.1:8080")
}
