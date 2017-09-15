package db

import "log"

func initTable(){
	initLogTable()
}

func initLogTable(){
	sql:=`create table if not exists logInfo(
		id SERIAL NOT NULL,
		level varchar(20),
		funcname	text,
		content text,
		createdat timestamp NOT NULL default now(),
		PRIMARY KEY ("id")
	);`
	stmt,err:=Db.Prepare(sql)
	if err != nil {
		log.Println("fail in initLogTable Prepare"+err.Error())
		return
	}
	_,err=stmt.Exec()
	defer stmt.Close()
	if err != nil {
		log.Println("fail in initLogTable stmt.Exec"+err.Error())
		return
	}
}
