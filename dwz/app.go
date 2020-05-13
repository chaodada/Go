package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/justinas/alice" // 方便的使用中间件

	"github.com/gorilla/mux"
	validate "gopkg.in/validator.v2" // 参数验证
)

// 核心数据结构
type App struct {
	Router      *mux.Router //
	Middlewares *Middleware //  中间件
	Config      *Env
}

// 请求结构体
type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

// 响应结构体
type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

// 初始化函数
func (a *App) Initialize(e *Env) {

	// 设置日志
	log.SetFlags(log.LstdFlags | log.Lshortfile) // log.LstdFlags  指定日志发生的时间与日期 | log.Lshortfile 行号与文件名

	a.Config = e
	fmt.Println("初始化日志成功")
	a.Router = mux.NewRouter()
	fmt.Println("初始化mux成功")
	a.Middlewares = &Middleware{}
	fmt.Println("初始化Middleware成功")
	a.initializeRoutes()
	fmt.Println("初始化路由成功")
}

// 初始化路由函数
func (a *App) initializeRoutes() {

	m := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler) // 初始化中间件

	//a.Router.HandleFunc("/", index).Methods("GET")
	//a.Router.HandleFunc("/post", post).Methods("POST")
	//a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")          //创建短网址
	//a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")             // 短网址详细信息
	//a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,8}}", a.redirect).Methods("GET") //重定向接口

	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortlink)).Methods("POST")          //创建短网址
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortlinkInfo)).Methods("GET")             // 短网址详细信息
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,8}}", m.ThenFunc(a.redirect)).Methods("GET") //重定向接口
}

//创建短网址
func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("请求参数解析错误%v", r.Body)})
		return
	}
	if err := validate.Validate(req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("请求参数验证错误%v", req)})
		return
	}
	defer r.Body.Close() // 在父函数返回之前执行

	s, err := a.Config.S.Shorten(req.URL, req.ExpirationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusCreated, shortlinkResp{Shortlink: s})
	}
	fmt.Printf("%v\n", req)
}

// 短网址详细信息
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func (a *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()      // 获取get的值
	s := vals.Get("shortlink") // 获取指定get参数

	d, err := a.Config.S.ShortlinkInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJSON(w, http.StatusOK, d)
	}

	//fmt.Printf("%s\n", s)
	//panic(s)
}

// 重定向接口
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {

	vals := mux.Vars(r)
	//fmt.Printf("%s\n", vals["shortlink"])
	url, err := a.Config.S.Unshorten(vals["shortlink"])
	if err != nil {
		respondWithError(w, err)
	} else {
		//respondWithJSON(w, http.StatusOK, d)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect) //临时重定向
	}
}

func (a *App) run(addr string) {
	//http.HandleFunc("/", a.Router)
	//http.Handle("/", a.Router)
	//err := http.ListenAndServe(addr, a.Router)
	//fmt.Println(4)
	//http.ListenAndServe(addr, a.Router)

	log.Fatal(http.ListenAndServe(addr, a.Router))

	// 设置路由，如果访问/，则调用index方法
	//http.HandleFunc("/", index)
	//// 启动web服务，监听9090端口
	//err := http.ListenAndServe(addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

}

func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e.Error())
		//fmt.Println(e.Error())
		//fmt.Println(reflect.TypeOf(e))
		respondWithJSON(w, e.Status(), e.Error())
	default:
		respondWithJSON(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	//fmt.Println(payload)
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json") // 设置返回头
	w.WriteHeader(code)                                // 设置状态码
	w.Write(resp)
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
//func index(w http.ResponseWriter, r *http.Request) {
//	// 往w里写入内容，就会在浏览器里输出
//	fmt.Fprintf(w, "Hello golang http! /")
//}
//
//// w表示response对象，返回给客户端的内容都在对象里处理
//// r表示客户端请求对象，包含了请求头，请求参数等等
//func post(w http.ResponseWriter, r *http.Request) {
//	// 往w里写入内容，就会在浏览器里输出
//	fmt.Fprintf(w, "Hello golang http! pots")
//}
