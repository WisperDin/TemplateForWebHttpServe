package controller

import (
	"net/http"
	"../model"
	"../common"
	."../constant"
	"../common/logger"
	"fmt"
)

func GetArticle(w http.ResponseWriter, r *http.Request){
	common.SetContent(w)

	//检查session
/*	errCode := common.CheckSession(r)
	if errCode!=SESSION_OK {
		common.ReturnFormat(w,errCode,nil)
		return
	}*/

	//empty safe
	limitStr:=handlePageParam(r)
	orderStr:=handleOrderParam(r)

	articles,err:=model.FindArticle("",limitStr,orderStr)
	if err != nil {
		common.ReturnFormat(w,ARTICLE_SYS_ERR,nil)
		logger.Error("GetArticle","FindArticle failed "+err.Error())
		return
	}
	if len(articles)<=0{
		common.ReturnFormat(w,ARTICLE_NOTEXIST,nil)
		return
	}
	common.ReturnFormat(w,ARTICLE_SUCCESS,articles)
	logger.Info("GetArticle",fmt.Sprintf("limitStr %s orderStr %s",limitStr,orderStr))
}
