package test

import (
	"fmt"
	"testing"

	"simple-douyin/controller"
)

func TestToken(t *testing.T) {
	token, err := controller.GenerateToken("user")
	if err != nil {
		t.Error("GenerateToken error")
	}
    
	fmt.Println("token:", token)
    
	username, err := controller.ParseToken(token)
	if err != nil {
		t.Error("ParseToken error")
	}
	if username != "user" {
		t.Errorf("expect username, but got %s", username)
	}
}