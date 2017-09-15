package model

import (
	"time"
	"fmt"
	."../db"
)

type LogInfo struct {
	ID          int64  `json:"id"`
	Level       string `json:"level"`
	FuncName 		string `json:"funcName"`
	Content     string `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

const LogInfoTableName = "loginfo"

func (l *LogInfo) Insert() (err error) {
	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO %s(level,funcname,content,createdat) "+
		"VALUES($1,$2,$3,$4)", LogInfoTableName))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	l.CreatedAt = time.Now()
	_, err = stmt.Exec(l.Level, l.FuncName,l.Content,l.CreatedAt)
	return
}
