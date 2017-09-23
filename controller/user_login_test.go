package controller

import (
	"testing"
)


func TestLogin(t *testing.T) {
	testInit()
	mockPostRqs(t, "/api/login", "username=123&pwd=1234","")
}
