package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice" // 方便的使用中间件
)

// web应用类
type App struct {
	Router      *mux.Router
	Middlewares *Middleware
}

// 应用初始化函数
func (a *App) Initalize() {

	a.Router = mux.NewRouter()
	log.Println("初始化mux成功")

	a.Middlewares = &Middleware{}
	log.Println("初始化Middleware成功")

	a.registerHandler()
	log.Println("初始化路由成功")
}

// 注册路由函数
func (a *App) registerHandler() {

	//不走中间件
	//a.Router.HandleFunc("/", index).Methods("GET")
	//a.Router.HandleFunc("/post", post).Methods("POST")
	//a.Router.HandleFunc("/get", get).Methods("GET")

	// 设置中间件
	chain := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)
	log.Println("设置中间件成功")

	//创建短网址
	a.Router.Handle("/api/shorten", chain.ThenFunc(createShortUrl)).Methods("POST")
	log.Println("创建短网址--路由设置成功")

	// 短网址详细信息
	a.Router.Handle("/api/info", chain.ThenFunc(getShortlinkInfo)).Methods("GET")
	log.Println("短网址详细信息--路由设置成功")

	//重定向接口
	a.Router.Handle("/{shortUrl:[a-zA-Z0-9]{1,11}}", chain.ThenFunc(redirect)).Methods("GET")
	log.Println("重定向接口--路由设置成功")

}

// 运行
func (a *App) Run(addr string) {
	// 启动web服务，监听9090端口
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		// 监听端口失败
		log.Printf("Server startup failed.")
		panic(err)
	}
}
