package controller

import (
	"testing"

	"../common"
	"../model"
	"log"
)

func initArticleData(t *testing.T){
	//初始数据
	article1:=new(model.Artcle)
	article1.Theme="主题1"
	article1.Content="内容1"
	article1.ImgUrl="1"
	err:=article1.Insert()
	if err != nil {
		t.Error(err)
		return
	}

	article2:=new(model.Artcle)
	article2.Theme="主题2"
	article2.Content="内容2"
	article2.ImgUrl="2"
	err=article2.Insert()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetArticle(t *testing.T) {

	initArticleData(t)

	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)

	mockGetRqs(t,"/api/article?username=123&userid=1",token,
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

func TestGetArticleEmpty(t *testing.T) {
	//模拟session
	userid:=1
	usernameRaw:="123"
	token:=common.SaveSession(int64(userid),usernameRaw)

	mockGetRqs(t,"/api/article?username=123&userid=1",token,
		func(rawBody string,r *common.R) (bool) {
			if r==nil {
				t.Error("r null")
				return false
			}
			if r.Code != 2005{
				log.Println(r)
				t.Error("r.Code != 2005")
				return false
			}
			return true
		})
	teardown()
}
//todo...分页 排序的功能测试
