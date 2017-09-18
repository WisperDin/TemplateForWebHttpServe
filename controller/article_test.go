package controller

import (
	"testing"

)

//todo 已有问题 有些列为空就scan出错

func TestGetArticle(t *testing.T) {
	testInit()
	mockGetArticle(t,"/api/article?username=123&userid=1",0)
}
