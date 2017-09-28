package model

import (
	"time"
	"fmt"
	."../db"

)

type User struct {
	ID          int64  `json:"id"`
	UserName       string `json:"username"`
	Password 		string `json:"-"`
	CreatedAt   time.Time `json:"createdat"`
}

const userTableName = "user"

func (a *User) Insert() (err error) {
	stmt, err := Db.Prepare(fmt.Sprintf(`INSERT INTO "%s"(UserName,Password,createdat) VALUES($1,$2,$3)`, userTableName))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	a.CreatedAt = time.Now()
	_, err = stmt.Exec(a.UserName, a.Password,a.CreatedAt)
	return
}

func FindUser(condition, limit, order string) ([]User, error) {
	result := []User{}
	rows, err := Db.Query(fmt.Sprintf(`SELECT id,username,password,createdat FROM "%s" %s %s %s`, userTableName, condition, order, limit))
	if err != nil {
		return result, err
	}
	for rows.Next() {
		tmp := User{}
		err = rows.Scan(&tmp.ID, &tmp.UserName, &tmp.Password, &tmp.CreatedAt)
		result = append(result, tmp)
	}
	return result, err
}
