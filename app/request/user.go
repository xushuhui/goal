package request

import (
	"github.com/gin-gonic/gin"
	"go-web/app/utils"
	"net/http"
)

type UserLoginRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserLogin(c *gin.Context) *UserLoginRequest {
	var user *UserLoginRequest
	err := utils.Bind(c.Request, &user)
	if err != nil {
		panic(err)
	}
	return user
}

type UserRegRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserReg(Request *http.Request) *UserRegRequest {
	var user *UserRegRequest
	err := utils.Bind(Request, &user)
	if err != nil {
		panic(err)
	}
	return user
}
