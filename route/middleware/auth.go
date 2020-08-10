package middleware

import (
	"github.com/gin-gonic/gin"
	"goal/core"
	"goal/lib"
	"goal/pkg/errcode"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			core.FailResp(c, errcode.InvalidParams)
			return
		}

		claims, err := lib.ParseToken(Authorization)
		if err != nil {
			core.FailResp(c, errcode.ErrorAuthToken)
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			core.FailResp(c, errcode.TimeoutAuthToken)
			return
		}

		c.Next()
	}

}
