package model

import (
	"time"
	"fmt"
	."../db"
	"database/sql"
)

type MsgBoard struct {
	ID          int64  `json:"id"`
	Content     string `json:"content"`
	authorid      sql.NullInt64 `json:"-"`
	Authorid      int64 `json:"authorid"`
	CreatedAtTime   time.Time `json:"-"`
	CreatedAt   string `json:"createdat"`
}

const msgBoardTableName  = "msgboard"

func (a *MsgBoard) Insert() (err error) {
	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO %s(content,authorid,createdat) "+
		"VALUES($1,$2,$3)", msgBoardTableName))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	a.CreatedAtTime = time.Now()
	_, err = stmt.Exec(a.Content, a.Authorid,a.CreatedAtTime)
	return
}

func FindMsg(condition, limit, order string) ([]*MsgBoard, error) {
	result := []*MsgBoard{}
	rows, err := Db.Query(fmt.Sprintf(`SELECT id,content,authorid,createdat FROM %s %s %s %s`, msgBoardTableName, condition, order, limit))
	if err != nil {
		return result, err
	}
	for rows.Next() {
		tmp := MsgBoard{}
		err = rows.Scan(&tmp.ID, &tmp.Content, &tmp.authorid,&tmp.CreatedAtTime)
		if tmp.authorid.Valid {
			tmp.Authorid=tmp.authorid.Int64
		}
		result = append(result, &tmp)
	}
	return result, err
}

