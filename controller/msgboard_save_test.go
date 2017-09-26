package controller

import (
	"testing"

	"../common"
)

func TestSaveMsg(t *testing.T) {
	testInit()
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)
	mockPostRqs(t,"/api/insert/msg","username=123&userid=1&content=嘿嘿",token)
}
