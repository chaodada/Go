package web

import (
	"duan/defs"
	"log"
	"net/http"
	"time"
)

// 中间件结构体
type Middleware struct {
}

// 记录请求需要消耗的时间
func (m Middleware) LoggingHandler(next http.Handler) http.Handler {
	// 匿名函数
	// w表示response对象，返回给客户端的内容都在对象里处理
	// r表示客户端请求对象，包含了请求头，请求参数等等
	fn := func(w http.ResponseWriter, r *http.Request) {
		// 开始时间
		startTime := time.Now()

		// 它要做的就是根据请求的数据做处理，把结果写到 Response 中
		next.ServeHTTP(w, r)

		// 结束时间
		endTime := time.Now()

		// 打印日志
		log.Printf("请求方式>>>>>>>[%s] 请求地址>>>>>>>[%q] 请求耗时>>>>>>>%v", r.Method, r.URL.String(), endTime.Sub(startTime))
	}
	return http.HandlerFunc(fn)
}

// 把程序从panic 恢复过来 (不懂)
func (m Middleware) RecoverHandler(next http.Handler) http.Handler {
	// 匿名函数
	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recover from panic:%+v", err)
				sendErrorResponse(w, defs.ErrorInternalFaults)
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
