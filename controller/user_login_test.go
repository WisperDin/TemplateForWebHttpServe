package controller

import (
	"testing"
)


func TestLogin(t *testing.T) {
	testInit()

	//todo.自动插入测试用户
	mockPostRqs(t, "/api/login", "username=123&pwd=1234","")
	//todo.清除用户数据
}
