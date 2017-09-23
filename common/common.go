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
	res := R{Code: code, Data: data}
	fb, _ := json.Marshal(res)
	w.Write(fb)
}

func SetContent(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4200")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
