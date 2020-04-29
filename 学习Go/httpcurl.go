package main

import "fmt"
import "net/http"
import "io/ioutil"
import "strings"
import "net/http/cookiejar"

func main() {
	resp, cookie := httpCurl("POST", "http://192.168.9.214/index/index/login", "username=admin&password=admin", nil)
	fmt.Printf("%s", resp)
	resp, cookie = httpCurl("GET", "http://192.168.9.214/index/index/index", "", cookie)
	fmt.Printf("%s", resp)
}

// httpCurl 模拟网络请求
// 参数:method 模式 ["GET","POST"] / url / data / cookie
// 返回:return body 响应内容 / 返回 cookie
func httpCurl(method string, url string, data string, cookie []*http.Cookie) ([]byte, []*http.Cookie) {
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, strings.NewReader(data))
	jar, _ := cookiejar.New(nil)
	jar.SetCookies(req.URL, cookie)
	client.Jar = jar
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	re_cookie := resp.Cookies()
	resp.Body.Close()
	return body, re_cookie
}
