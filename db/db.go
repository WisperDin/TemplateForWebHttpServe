package db

import (
	"fmt"
	"database/sql"
	//加对应database的lib
	_ "github.com/lib/pq"
	"log"
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
		log.Println("InitDB failed at Ping "+err.Error())
		panic(err)
		return err
	}
	//初始化table
	initAllTable()
	return nil
}




