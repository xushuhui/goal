/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package routes

import (
	"github.com/gin-gonic/gin"
	"go-web/app/controller"
	"go-web/app/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.Statcost())
	router.POST("/user/login", controller.UserLogin)
	router.POST("/user/register", controller.UserReg)
	router.Use(middleware.AuthToken())
	router.GET("/book", controller.BookList)
	return router

}
