package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/core"
	"go-web/core/ext"
	"net/http"
)

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		claims, err := ext.ParseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, core.Fail(core.SYSTEMERROR))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
