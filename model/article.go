package model

import (
	"time"
	"fmt"
	."../db"
)

type Artcle struct {
	ID          int64  `json:"id"`
	Theme       string `json:"theme"`
	Content     string `json:"content"`
	ImgUrl      string `json:"imgurl"`
	CreatedAt   time.Time `json:"createdat"`
	UpdateAt    time.Time `json:"updateat"`
}

const articleTableName  = "article"

//todo 已有问题 有些列为空就scan出错

func FindArticle(condition, limit, order string) ([]Artcle, error) {
	result := []Artcle{}
	rows, err := Db.Query(fmt.Sprintf(`SELECT id,theme,content,imgurl,createdat,updateat FROM %s %s %s %s`, articleTableName, condition, order, limit))
	if err != nil {
		return result, err
	}
	for rows.Next() {
		tmp := Artcle{}
		err = rows.Scan(&tmp.ID, &tmp.Theme, &tmp.Content, &tmp.ImgUrl,&tmp.CreatedAt,&tmp.UpdateAt)
		result = append(result, tmp)
	}
	return result, err
}

