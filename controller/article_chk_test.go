package controller

import (
	"testing"

	"../common"
)


func TestGetArticle(t *testing.T) {
	testInit()
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)

	mockGetRqs(t,"/api/article?username=123&userid=1",token)
}

//todo...分页 排序的功能测试
