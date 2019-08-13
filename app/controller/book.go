package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web/app/request"
	"go-web/app/services"
	"go-web/core"
)

func BookList(c *gin.Context) {
	claims, _ := c.Get("claims")
	fmt.Println(claims)
}

func Read(c *gin.Context) {
	if c.Request == nil {
		core.JsonResp(c, core.Fail(core.SYSTEMERROR))
		return
	}
	resp := services.UserLogin(request.UserLogin(c))
	core.JsonResp(c, resp)
}
