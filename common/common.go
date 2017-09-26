package common

import (
	"net/http"
	"encoding/json"
	"../conf"
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
	w.Header().Set("Access-Control-Allow-Origin", conf.App.WebPageHost)             //允许访问所有域 	//允许访问所有域

	w.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}
