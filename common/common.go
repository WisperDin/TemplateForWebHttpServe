package common

import (
	"net/http"
	"encoding/json"
)

type R struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ReturnFormat(w http.ResponseWriter, code int, data interface{}) {
	SetContent(w)
	res := R{Code: code, Data: data}
	fb, _ := json.Marshal(res)
	w.Write(fb)
}

func SetContent(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
}
