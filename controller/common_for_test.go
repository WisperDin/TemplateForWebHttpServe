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
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBTestName,conf.App.DBDriver)
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
	r.HandleFunc("/api/insert/article",
		SaveArticle).
		Methods(http.MethodPost)
	r.HandleFunc("/api/insert/msg",
		SaveMsg).
		Methods(http.MethodPost)
	r.HandleFunc("/api/msg",
		GetMsg).
		Methods(http.MethodGet)


	return r
}



func mockGetArticle(t *testing.T, url string, expectedCode int, token string){
/*	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"*/

	cookie := make(map[string]string)
	cookie["Token"] = token

	r.GET(url).
		SetDebug(true).
		SetCookie(cookie).
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
		log.Println(string(body))
	})
}

func mockGetRqs(t *testing.T, url string,token string) {
	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"

	cookie := make(map[string]string)
	cookie["Token"] = token
	log.Println(r)
	r.GET(url).
		SetDebug(true).
		SetHeader(headers).
		SetCookie(cookie).
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
		log.Println(string(body))
	})
}


func mockPostRqs(t *testing.T, url string, body string,token string) {
	headers := make(map[string]string)
	headers["Content-type"] = "application/x-www-form-urlencoded"

	cookie := make(map[string]string)
	cookie["Token"] = token
	log.Println(r)
	r.POST(url).
		SetDebug(true).
		SetHeader(headers).
		SetBody(body).
		SetCookie(cookie).
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

