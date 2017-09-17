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
	"./view"
	"os"
	"strconv"
	"io"
)

func main() {
	conf.Init("./app.toml")
	db.InitDB(conf.App.DBHost, conf.App.DBPort, conf.App.DBUser, conf.App.DBPassword, conf.App.DBName,conf.App.DBDriver)
	log.Println(conf.App)
	logger.Init()
	view.Init("./dist/")

	r := mux.NewRouter()

	//serve frontend artifact
	r.PathPrefix("/page/").HandlerFunc(view.LoadTemplate)

	//api
	r.HandleFunc("/exit", httpShutDownHandler)

	r.HandleFunc("/api/login",controller.Login).Methods(http.MethodPost,http.MethodOptions)

	//对于/路由 要放后面
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./dist/"))))
	http.ListenAndServe(fmt.Sprintf(":%s", conf.App.ServerPort), r)

}

//关闭当前进程
func httpShutDownHandler(w http.ResponseWriter, r *http.Request) {
	shutdownMsg := "\n\n lzyweb shutdown\n\n"

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(shutdownMsg)))
	io.WriteString(w, shutdownMsg)

	f, canFlush := w.(http.Flusher)
	if canFlush {
		f.Flush()
	}

	conn, _, err := w.(http.Hijacker).Hijack()
	if err != nil {
		logger.Error("httpShutDownHandler","Hijack failed")
		return
	}

	conn.Close()

	logger.Info("httpShutDownHandler",shutdownMsg)
	os.Exit(0)
}
