/**
 * Session模块
 *
 */
package common

import (
	"fmt"
	//"net/http"
	"./logger"
	"strconv"
	"crypto/md5"
	"time"
	"io"
)

type Session struct {
	UserId 	  int64
	UserName  string
	Token     string
}

var UserSession map[string]*Session

func GetSessionID(id int64,username string) string {
	if id<=0||len(username)<=0{
		logger.Error("GetSessionID","param error")
		return ""
	}
	return strconv.Itoa(int(id))+"-"+username
}

//保存一个session
func SaveSession(id int64,username string) (token string) {
	if id<=0||len(username)<=0{
		logger.Error("SaveSession","param error")
		return
	}

	if UserSession == nil {
		UserSession = map[string]*Session{}
	}
	sessionKey := GetSessionID(id,username)
	session := new(Session)
	session.UserId=id
	session.UserName=username
	//计算session签名
	UserSession[sessionKey] = session

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token = fmt.Sprintf("%x", h.Sum(nil))
	session.Token = token
	return
}

/*//移除一个session
func RemoveSession(r *http.Request) {
*//*	sessionKey := GetSessionID()
	session := new(Session)
	if UserSession == nil {
		UserSession = map[string]*Session{}
	}
	UserSession[sessionKey] = session*//*
}

//检查session是否合法
func CheckSession(r *http.Request) error {
*//*	sessionKey := GetSessionID()
	session := UserSession[sessionKey]
	if session == nil {
		err := fmt.Errorf("用户session校验失败，session不存在！")
		return err
	}
	return nil*//*
}*/
