package controller

import (
	"testing"
	"../constant"


)


func TestLogin(t *testing.T) {
	testInit()
	mockLoginRqs(t, "/api/login", constant.LOGIN_SUCCESS)
}
