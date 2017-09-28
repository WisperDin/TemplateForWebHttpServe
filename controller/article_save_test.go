package controller

import (
	"testing"

	"../common"
)

func TestSaveArticle(t *testing.T) {
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)
	mockPostRqs(t,"/api/insert/article","username=123&userid=1&theme=测试&content=嘿嘿&imgurl=abcd",token,
		func(rawBody string,r *common.R) (bool) {
		if r==nil {
			t.Error("r null")
			return false
		}
		if r.Code != 2000{
			t.Error("r.Code != 2000")
			return false
		}
		return true
	})
	teardown()
}
