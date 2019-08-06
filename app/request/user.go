package request

import (
	"github.com/gin-gonic/gin"
	"go-web/app/utils"
)

type UserLoginRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserLogin(c *gin.Context) *UserLoginRequest {
	var user *UserLoginRequest
	err := utils.BindJson(c, &user)
	if err != nil {
		panic(err)
	}
	return user
}

type UserRegRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserReg(c *gin.Context) *UserRegRequest {
	var user *UserRegRequest
	err := utils.BindJson(c, &user)
	if err != nil {
		panic(err)
	}
	return user
}
