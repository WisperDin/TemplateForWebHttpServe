package model

import (
	"time"
	"fmt"
	."../db"
	"database/sql"
)

type Artcle struct {
	ID          int64  `json:"id"`
	Theme       string `json:"theme"`
	Content     string `json:"content"`
	imgUrl      sql.NullString `json:"-"`
	ImgUrl      string `json:"imgurl"`
	CreatedAtTime   time.Time `json:"-"`
	UpdateAt    time.Time `json:"updateat"`
	CreatedAt   string `json:"createdat"`
}

const articleTableName  = "article"

func (a *Artcle) Insert() (err error) {
	stmt, err := Db.Prepare(fmt.Sprintf("INSERT INTO %s(theme,content,imgurl,createdat,updateat) "+
		"VALUES($1,$2,$3,$4,$5)", articleTableName))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	a.CreatedAtTime = time.Now()
	a.UpdateAt = a.CreatedAtTime
	_, err = stmt.Exec(a.Theme, a.Content,a.ImgUrl,a.CreatedAtTime,a.UpdateAt)
	return
}

func FindArticle(condition, limit, order string) ([]*Artcle, error) {
	result := []*Artcle{}
	rows, err := Db.Query(fmt.Sprintf(`SELECT id,theme,content,imgurl,createdat,updateat FROM %s %s %s %s`, articleTableName, condition, order, limit))
	if err != nil {
		return result, err
	}
	for rows.Next() {
		tmp := Artcle{}
		err = rows.Scan(&tmp.ID, &tmp.Theme, &tmp.Content, &tmp.imgUrl,&tmp.CreatedAtTime,&tmp.UpdateAt)
		if tmp.imgUrl.Valid {
			tmp.ImgUrl=tmp.imgUrl.String
		}
		result = append(result, &tmp)
	}
	return result, err
}

