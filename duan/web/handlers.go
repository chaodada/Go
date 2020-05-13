package web

import (
	"duan/conf"
	"duan/defs"
	"duan/storage"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

var s storage.Storage

func init() {
	redis := conf.C.Redis
	s = storage.NewRedisCli(redis.Addr, redis.Pwd, redis.Db)
}

// 创建短网址
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func createShortUrl(w http.ResponseWriter, r *http.Request) {

	// 拿到请求结构体
	req := &defs.ShortenReq{}

	fmt.Println(r.Body)
	fmt.Println(r.Header)

	// 从请求body中读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	// Decode从输入流读取下一个json编码值并保存在v指向的值里，
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
		return
	}
	defer r.Body.Close()

	s, err := s.Shorten(req.Url, req.ExpirationInMinutes)
	if err != nil {
		sendErrorResponse(w, defs.ErrorStorageError)
		return
	}
	sendNormalResponse(w, defs.Result{Code: "0", ErrMsg: "OK", Data: defs.ShortenResp{ShortUrl: s, LongUrl: req.Url}}, http.StatusOK)
}

func getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shortUrl")
	info, err := s.ShortlinkInfo(shortUrl)
	if err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		urlDetail := &storage.URLDetail{}
		_ = json.Unmarshal([]byte(info.(string)), urlDetail)
		sendNormalResponse(w, defs.Result{Code: "0", ErrMsg: "OK", Data: urlDetail}, http.StatusOK)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {

	fmt.Println("sdsadsada")

	vars := mux.Vars(r)
	fmt.Println(vars)
	shortUrl := vars["shortUrl"]
	longUrl, err := s.Unshorten(shortUrl)
	if err != nil {
		sendErrorResponse(w, defs.Response{HttpSC: http.StatusInternalServerError, Result: defs.Result{Code: "007", ErrMsg: err.Error()}})
	} else {
		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	}
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "测试get / ")
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func post(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "测试post /post")
}

// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func get(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "测试get /get")
}
