package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/app/request"
	"go-web/app/service"
	"go-web/core"
)

//UserLogin func
func UserLogin(c *gin.Context) {
	if c.Request == nil {
		core.JsonResp(c, core.Fail(core.SYSTEMERROR))
		return
	}
	resp := service.UserLogin(request.UserLogin(c))
	core.JsonResp(c, resp)
}
func UserReg(c *gin.Context) {
	//userRequest := request.UserReg(r)
	//service.UserReg(userRequest)
}
