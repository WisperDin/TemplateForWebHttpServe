package main

import (
	"./conf"
	"log"
	"./db"
	"fmt"
	"net/http"
	"./common/logger"
	"github.com/gorilla/mux"
	"./controller"
)

func main() {
	conf.Init("./app.toml")
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName,conf.App.DBDriver)
	log.Println(conf.App)
	logger.Init()

	//serve frontend artifact
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./dist"))))

	//api
	r := mux.NewRouter()

	r.HandleFunc("/api/login",controller.Login).Methods(http.MethodPost)

	http.ListenAndServe(fmt.Sprintf(":%s", conf.App.ServerPort), r)

}
