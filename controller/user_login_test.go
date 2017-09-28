package controller

import (
	"testing"
	"../common"
	"log"
	"../model"
)

func initUserData(t *testing.T){
	user:=new(model.User)
	user.UserName="testN"
	user.Password="123"
	err:=user.Insert()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestLogin(t *testing.T) {

	initUserData(t)

	mockPostRqs(t, "/api/login", "username=testN&pwd=123","",
		func(rawBody string,r *common.R) (bool) {
		if r==nil {
			t.Error("r null")
			return false
		}
		if r.Code != 1000{
			t.Error("r.Code != 1000")
			return false
		}
		log.Println(r)
		return true
	})
	teardown()
}

func TestLoginFail(t *testing.T) {

	initUserData(t)

	mockPostRqs(t, "/api/login", "username=4123&pwd=123","",
		func(rawBody string,r *common.R) (bool) {
			if r==nil {
				t.Error("r null")
				return false
			}
			if r.Code != 1005{
				t.Error("r.Code != 1005")
				return false
			}
			log.Println(r)
			return true
		})
	teardown()
}
