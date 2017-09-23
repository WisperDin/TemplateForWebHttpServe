/**
 * Session模块
 *
 */
package common

import (
	"fmt"
	"./logger"
	"strconv"
	"crypto/md5"
	"time"
	"io"
	"net/http"
	."../constant"
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

//检查session是否合法
func CheckSession(r *http.Request) int {
	useridRaw := r.FormValue("userid")
	usernameRaw := r.FormValue("username")
	if len(useridRaw)<=0||len(usernameRaw)<=0{
		return SESSION_PARA_ERR
	}
	userid,err:=strconv.Atoi(useridRaw)
	if err != nil {
		return SESSION_SYS_ERR
	}

	sessionKey := GetSessionID(int64(userid),usernameRaw)
	session := UserSession[sessionKey]
	if session == nil {
		return SESSION_NO_AUTH
	}

	//从cookie中获取token
	cookie,err:=r.Cookie("token")
	if err != nil {
		logger.Error("CheckSession","get Cookie failed "+err.Error())
		return SESSION_NO_COOKIE
	}
	if cookie.Value!=session.Token {
		return SESSION_TOKEN_NOTMATCH
	}

	return SESSION_OK
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
*/


