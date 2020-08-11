package api

import (
	"github.com/gin-gonic/gin"
	"goal/internal/request"
	"goal/internal/services"
	"goal/pkg/core"
)

//UserLogin func
func Login(c *gin.Context) {
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
