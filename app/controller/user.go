package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"go-web/core"
)

//UserLogin func
func UserLogin(c *gin.Context) {
	if c.Request == nil {
		core.JsonResp(c, core.Fail(core.SYSTEMERROR))
		return
	}
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))

	//resp := service.UserLogin(request.UserLogin(c))
	//core.JsonResp(c, resp)
}
func UserReg(c *gin.Context) {
	//userRequest := request.UserReg(r)
	//service.UserReg(userRequest)
}
