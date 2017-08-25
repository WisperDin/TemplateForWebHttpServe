/**
 * Session模块
 *
 */
package common

import (
	"fmt"
	"net/http"
)

type Session struct {
}

var UserSession map[string]*Session

func GetSessionID() string {
	return ""
}

//保存一个session
func SaveSession() (token string) {
	if UserSession == nil {
		UserSession = map[string]*Session{}
	}
	sessionKey := GetSessionID()
	session := new(Session)
	//计算session签名
	UserSession[sessionKey] = session
	return
}

//移除一个session
func RemoveSession(r *http.Request) {
	sessionKey := GetSessionID()
	session := new(Session)
	if UserSession == nil {
		UserSession = map[string]*Session{}
	}
	UserSession[sessionKey] = session
}

//检查session是否合法
func CheckSession(r *http.Request) error {
	sessionKey := GetSessionID()
	session := UserSession[sessionKey]
	if session == nil {
		err := fmt.Errorf("用户session校验失败，session不存在！")
		return err
	}
	return nil
}
