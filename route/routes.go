package route

import (
	"github.com/gin-gonic/gin"
	"goal/app/api"
	"goal/route/middleware"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors(), middleware.Logger(), middleware.ErrorHandle())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	router.POST("/login", api.Login)
	router.Use(middleware.Auth())
	//需要登录的接口
	return router
}
