package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/core"
	"net/http"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("bear ")
		if len(token) < 3 {
			c.Abort()
			c.JSON(http.StatusOK, core.Fail(core.SYSTEMERROR))
		}
		c.Next()
	}
}
