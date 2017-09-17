package controller

import (
	"net/http"
	"../common"
	"../model"
	"fmt"
	."../constant"
	"../common/logger"
)

func Login(w http.ResponseWriter, r *http.Request){
	common.SetContent(w)
	username := r.PostFormValue("username")
	pwd := r.PostFormValue("pwd")
	if username == "" {
		common.ReturnFormat(w, LOGIN_PARA_ERR, nil)
		return
	}
	if pwd == "" {
		common.ReturnFormat(w, LOGIN_PARA_ERR, nil)
		return
	}

	//查用户
	user,err:=model.FindUser(fmt.Sprintf("where username='%s'",username),"","")
	if err != nil {
		logger.Error("controller-Login",fmt.Sprintf("User %s model.FindUser failed ",username)+err.Error())
		common.ReturnFormat(w, LOGIN_SYS_ERR, nil)
		return
	}

	if len(user)<=0 {
		common.ReturnFormat(w, LOGIN_USER_NOTEXIST, nil)
		return
	}

	//登录失败
	if user[0].Password!=pwd{
		common.ReturnFormat(w, LOGIN_WRONG_PWD, nil)
		return
	}

	logger.Info("controller-Login",fmt.Sprintf("User %s login success",username))

	//设置session
	token := common.SaveSession(user[0].ID,user[0].UserName)
	cookie := http.Cookie{Name: "Token", Value: token, Path: "/", MaxAge: 86400}
	http.SetCookie(w,&cookie)
	common.ReturnFormat(w, LOGIN_SUCCESS, user[0])
	return
}
