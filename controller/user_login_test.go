package controller

import (
	"testing"
	"../conf"
	"../common"
	"../common/logger"
	"../db"

	"github.com/gorilla/mux"
	"net/http"
	"github.com/appleboy/gofight"

	"encoding/json"
	//"github.com/magiconair/properties/assert"
	"log"
	"../constant"
	"io/ioutil"

)

func init()  {
	conf.Init("../app.toml")
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName,conf.App.DBDriver)
	logger.Init()

	r = gofight.New()
}

var r *gofight.RequestConfig

func MuxEngine() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/login",
		Login).
		Methods(http.MethodPost)
	return r
}

func mockLoginRqs(t *testing.T, url string, expectedCode int) {
	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"

	body:="username=123&pwd=1234"
	r.POST(url).
		SetDebug(true).
		SetHeader(headers).
		SetBody(body).
		Run(MuxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		fb := &common.R{}
		err = json.Unmarshal(body, fb)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(fb)
		//assert.Equal(t, expectedCode, fb.Code)
	})
}


func TestLogin(t *testing.T) {
	mockLoginRqs(t, "/api/login", constant.LOGIN_SUCCESS)
}
