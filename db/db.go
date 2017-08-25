package db

import (
	"fmt"
	"database/sql"
	//加对应database的lib
	_ "github.com/lib/pq"
	"time"
)

var Db *sql.DB
var dBDriverName string

func InitDB(host, port, user, pwd, dbName, driverName string) error {
	dBDriverName = driverName
	//构建连接字符串
	dateSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pwd, dbName)
	fmt.Println(dateSource)
	db, _ := sql.Open(dBDriverName, dateSource)
	Db = db
	err := Db.Ping()
	if err != nil {
		//自动重连
		go reInit(dateSource)
	}
	return nil
}

func reInit(dateSource string) {
	for {
		db, _ := sql.Open(dBDriverName, dateSource)
		if err := db.Ping(); err == nil {
			Db = db
			break
		} else {
			fmt.Println("数据库连接失败，2分钟后重试")
			time.Sleep(time.Minute * 2)
			reInit(dateSource)
		}
	}
}
