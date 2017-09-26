package controller

import (
	"../common"
	"net/http"
	"../model"
	."../constant"
	"../common/logger"
	"fmt"
)

func SaveMsg(w http.ResponseWriter, r *http.Request){
	common.SetContent(w)

	//todo 检查session

	content := r.FormValue("content")
	if len(content)<=0{
		common.ReturnFormat(w,MSG_PARA_ERR,nil)
		return
	}

	msg:=new(model.MsgBoard)
	msg.Content = content
	err:=msg.Insert()
	if err != nil {
		logger.Error("SaveMsg","msg.Insert failed "+err.Error())
		common.ReturnFormat(w,MSG_SYS_ERR,nil)
		return
	}

	common.ReturnFormat(w,MSG_SUCCESS,nil)
	logger.Info("SaveMsg","SaveMsg SUCCESS by user XX")
}

func GetMsg(w http.ResponseWriter, r *http.Request){
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

	msgs,err:=model.FindMsg("",limitStr,orderStr)
	if err != nil {
		common.ReturnFormat(w,MSG_SYS_ERR,nil)
		logger.Error("GetMsg","FindMsgBoard failed "+err.Error())
		return
	}
	if len(msgs)<=0{
		common.ReturnFormat(w,MSG_NOTEXIST,nil)
		return
	}

	//time format
	for _,msg:=range msgs {
		(*msg).CreatedAt=(*msg).CreatedAtTime.Format("2006-01-02 15:04:05")
	}

	common.ReturnFormat(w,MSG_SUCCESS,msgs)
	logger.Info("GetMsg",fmt.Sprintf("limitStr %s orderStr %s by user XX",limitStr,orderStr))
}