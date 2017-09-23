package controller

import (
	"net/http"
	"../model"
	"../common"
	."../constant"
	"../common/logger"
	"fmt"
)

func SaveArticle(w http.ResponseWriter, r *http.Request){
	common.SetContent(w)

	//检查session
/*	errCode := common.CheckSession(r)
	if errCode!=SESSION_OK {
		common.ReturnFormat(w,errCode,nil)
		return
	}*/

	theme := r.FormValue("theme")
	content := r.FormValue("content")
	imgurl := r.FormValue("imgurl")

	if len(theme)<=0{
		common.ReturnFormat(w,ARTICLE_PARA_ERR,nil)
		return
	}

	if len(content)<=0{
		common.ReturnFormat(w,ARTICLE_PARA_ERR,nil)
		return
	}

	//todo imgurl 判断

	article:=new(model.Artcle)
	article.Theme=theme
	article.Content=content
	article.ImgUrl=imgurl

	err:=article.Insert()
	if err != nil {
		logger.Error("SaveArticle","Insert failed "+err.Error())
		common.ReturnFormat(w,ARTICLE_SYS_ERR,nil)
		return
	}
	common.ReturnFormat(w,ARTICLE_SUCCESS,nil)
}


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
