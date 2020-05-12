package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	validate "gopkg.in/validator.v2"
)

type App struct {
	Router *mux.Router //
}

// 请求结构体
type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"expiration_in_minutes" validate:"min=0"`
}

func (a *App) Initialize() {

	//r := mux.NewRouter()

	r := mux.NewRouter()
	//r.HandleFunc("api/shorten", a.createShortlink).Methods("POST")

	//r.HandlerFunc("api/shorten", a.createShortlink).Methods("POST") //创建短网址
	a.Router = r
	//r.HandlerFunc("api/shorten", a.createShortlink).Methods("POST") //创建短网址
	//fmt.Println(r)
	//a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	//fmt.Println(a.Router)
	r := *a.Router
	r.HandleFunc("api/shorten", a.createShortlink).Methods("POST")           //创建短网址
	r.HandleFunc("api/info", a.getShortlinkInfo).Methods("GET")              // 短网址详细信息
	r.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,8}}", a.redirect).Methods("GET") //重定向接口
}

//创建短网址
func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}
	if err := validate.Validate(req); err != nil {
		return
	}
	defer r.Body.Close() // 在父函数返回之前执行
	fmt.Printf("%v\n", req)
}

// 短网址详细信息
func (a *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")
	fmt.Printf("%v\n", s)
}

//重定向接口
func (a *App) redirect(w http.ResponseWriter, r *http.Request) {

	vals := mux.Vars(r)
	fmt.Printf("%s\n", vals["shortlink"])
}
func (a *App) run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}
	a.Initialize()
	a.run(":8000")

}
