package middleware

import (
	"github.com/gin-gonic/gin"
	"go-web/core"
	"net/http"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.FormValue("token")
		if len(token) < 3 {
			c.Abort()

			c.JSON(http.StatusOK, core.Fail(core.SYSTEMERROR))

		}
	}
}
