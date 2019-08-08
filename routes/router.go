package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/controller"
	"go-web/app/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.POST("/user/login", controller.UserLogin)
	router.POST("/user/register", controller.UserReg)
	router.Use(middleware.AuthToken())
	router.GET("/book", controller.BookList)
	return router

}
