package main

import (
	"log"
	"net/http"
	"time"
)

// 定义中间件结构体
type Middleware struct {
}

// 记录请求需要消耗的时间
func (m Middleware) LoggingHandler(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()     // 开始时间
		next.ServeHTTP(w, r) // 对http请求的一个相应
		t2 := time.Now()     // 结束时间
		log.Printf("请求方式：[%s] 接口：%s 耗时： %v", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn) // 适配器 把用户自定义的函数 转换为
}

// 把程序从panic 恢复过来
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
