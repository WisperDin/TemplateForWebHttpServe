package db

import (
	"log"
	"errors"
)

func initAllTable(){
	var err error
	err = initTable(`create table if not exists logInfo(
		id SERIAL NOT NULL,
		level varchar(20),
		funcname	text,
		content text,
		createdat timestamp NOT NULL default now(),
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Println("initTable logInfo fail! "+err.Error())
		return
	}

	err = initTable(`create table if not exists "user"(
		id SERIAL NOT NULL,
		username varchar(20),
		password varchar(20),
		createdat timestamp NOT NULL default now(),
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Println("initTable user fail! "+err.Error())
		return
	}

	err = initTable(`create table if not exists article(
		id SERIAL NOT NULL,
		theme varchar(20) NOT NULL,
		content text NOT NULL,
		imgurl varchar(50),
		createdat timestamp NOT NULL default now(),
		updateat timestamp NOT NULL default now(),
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Println("initTable article fail! "+err.Error())
		return
	}

	err = initTable(`create table if not exists msgboard(
		id SERIAL NOT NULL,
		content text NOT NULL,
		authorid integer,
		createdat timestamp NOT NULL default now(),
		PRIMARY KEY ("id")
	);`)
	if err != nil {
		log.Println("initTable msgboard fail! "+err.Error())
		return
	}
}

func initTable(sql string)(error){
	if len(sql)<=0 {
		return errors.New("initTable sql empty")
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
