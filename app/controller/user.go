package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/app/request"
	"go-web/app/service"
	"go-web/core"
	"net/http"
)

//UserLogin func
func UserLogin(c *gin.Context) {
	if c.Request == nil {
		//http.Error(writer, "Please send a request body", 400)
		core.Fail(core.SYSTEMERROR)
		c.JSON(http.StatusOK, core.Fail(core.SYSTEMERROR))
		return
	}
	user := request.UserLogin(c.Request)
	service.UserLogin(writer, user)
}
func UserReg(writer http.ResponseWriter, r *http.Request) {
	userRequest := request.UserReg(r)
	service.UserReg(writer, userRequest)
}
