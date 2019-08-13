package request

import (
	"github.com/gin-gonic/gin"
	"go-web/app/utils"
)

type UserLoginRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func UserLogin(c *gin.Context) *UserLoginRequest {
	var request *UserLoginRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}

type UserRegRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func UserReg(c *gin.Context) *UserRegRequest {
	var request *UserRegRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}
