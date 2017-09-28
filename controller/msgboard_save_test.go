package controller

import (
	"testing"

	"../common"
	"log"
)

func TestSaveMsg(t *testing.T) {
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)
	mockPostRqs(t,"/api/insert/msg","username=123&userid=1&content=嘿嘿",token,
		func(rawBody string,r *common.R) (bool) {
			if r==nil {
				t.Error("r null")
				return false
			}
			if r.Code != 3000{
				t.Error("r.Code != 3000")
				return false
			}
			log.Println(r)
			return true
		})
	teardown()
}
