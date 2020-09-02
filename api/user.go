package api

import (
	"goal/global"
	"goal/internal/request"
	"goal/internal/services"
	"goal/pkg/core"

	"github.com/gin-gonic/gin"
)

/*

登录

URL:
	/login

POST参数：

	"phone":"2" //手机号
	"password": "1" //密码


返回值：

	"code": 0
	"message": "ok"

*/
func Login(c *gin.Context) {

	global.Logger.Errorf(c, "test")
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	err := services.Login(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	return

}
