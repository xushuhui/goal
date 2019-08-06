package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/controller"
	"go-web/app/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/user/login", controller.UserLogin)
	router.POST("/user/register", controller.UserReg)
	router.Use(middleware.VerifyToken())
	return router

}
