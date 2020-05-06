package real

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type URetriever struct { // 结构体
	UserAgent string
	TimeOut   time.Duration // 一个时间段
}

func (r *URetriever) Print() {
	fmt.Print(r.TimeOut, "    \n")
}

func (r *URetriever) Get(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		panic(err) // 报错 让程序停下来
	}
	result, err := httputil.DumpResponse(resp, true) //
	resp.Body.Close()                                //关闭
	if err != nil {
		panic(err)
	}
	return string(result)

}
