package web

import (
	"duan/defs"
	"encoding/json"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.Response) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8") //增加返回头
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(errResp.Result)
	_, _ = io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, result defs.Result, sc int) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8") //增加返回头
	w.WriteHeader(sc)
	bytes, _ := json.Marshal(result)
	_, _ = io.WriteString(w, string(bytes))
}
