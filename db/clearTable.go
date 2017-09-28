package db

import (
	"errors"
	"log"
)

//清空表!!
func ClearAllTable(){
	var err error
	err = clearTable(`delete from logInfo;`)
	if err != nil {
		log.Println("initTable logInfo fail! "+err.Error())
		return
	}

	err = clearTable(`delete from "user";`)
	if err != nil {
		log.Println("initTable user fail! "+err.Error())
		return
	}

	err = clearTable(`delete from article;`)
	if err != nil {
		log.Println("initTable article fail! "+err.Error())
		return
	}

	err = clearTable(`delete from msgboard;`)
	if err != nil {
		log.Println("initTable msgboard fail! "+err.Error())
		return
	}
	log.Println("clear table")
}

func clearTable(sql string)(error){
	if len(sql)<=0 {
		return errors.New("clearTable sql empty")
	}
	stmt,err:=Db.Prepare(sql)
	if err != nil {
		return err
	}
	_,err=stmt.Exec()
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

