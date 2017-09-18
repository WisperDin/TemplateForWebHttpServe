package controller

import (
	"../common"
	"github.com/appleboy/gofight"
	"io/ioutil"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
	"../conf"
	"../common/logger"
	"../db"
	"log"
	"encoding/json"
)

var r *gofight.RequestConfig

func testInit(){
	conf.Init("../app.toml")
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName,conf.App.DBDriver)
	logger.Init()
	r = gofight.New()
}

func muxEngine() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/login",
		Login).
		Methods(http.MethodPost)
	r.HandleFunc("/api/article",
		GetArticle).
		Methods(http.MethodGet)
	return r
}

func mockGetArticle(t *testing.T, url string, expectedCode int){
/*	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"*/
	r.GET(url).
		SetDebug(true).
		//SetHeader(headers).
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
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
	})
}

func mockLoginRqs(t *testing.T, url string, expectedCode int) {
	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"

	body:="username=123&pwd=1234"
	r.POST(url).
		SetDebug(true).
		SetHeader(headers).
		SetBody(body).
		Run(muxEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
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
	})
}
