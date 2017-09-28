package controller

import (
	"testing"

	"../common"
	"log"
	"../model"
)

func initMsgData(t *testing.T){
	msg1:=new(model.MsgBoard)
	msg1.Content = "msg1"
	err:=msg1.Insert()
	if err != nil {
		t.Error(err)
		return
	}

	msg2:=new(model.MsgBoard)
	msg2.Content = "msg2"
	err=msg2.Insert()
	if err != nil {
		t.Error(err)
		return
	}

}


func TestGetMsg(t *testing.T) {

	initMsgData(t)

	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)

	mockGetRqs(t,"/api/msg?username=123&userid=1",token,
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

func TestGetMsgEmpty(t *testing.T) {
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)

	mockGetRqs(t,"/api/msg?username=123&userid=1",token,
		func(rawBody string,r *common.R) (bool) {
			if r==nil {
				t.Error("r null")
				return false
			}
			if r.Code != 3005{
				t.Error("r.Code != 3005")
				return false
			}
			log.Println(r)
			return true
		})
	teardown()
}

//todo...分页 排序的功能测试
