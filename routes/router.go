/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package routes

import (
	"goal/api"
	"goal/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(middleware.StatCost())
	router.POST("/user/login", api.UserLogin)
	router.POST("/user/register", api.UserReg)
	router.Use(middleware.Auth())
	router.GET("/book", api.BookList)
	return router

}
