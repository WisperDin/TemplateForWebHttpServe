package main

import (
	"./conf"
	"log"
	"./db"
	"fmt"
	"net/http"
	"./common/logger"

)

func main() {
	conf.Init("./app.toml")
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName,conf.App.DBDriver)
	log.Println(conf.App)
	logger.Init()

	//api
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"welcome to my web server template!")
	})

	http.ListenAndServe(fmt.Sprintf(":%s", conf.App.ServerPort), nil)

}
