package middleware

import (
	"github.com/gin-gonic/gin"
	cd "goal/code"
	"goal/core"
	"goal/utils"
	"time"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		Authorization := c.Request.Header.Get("Authorization")

		if Authorization == "" {
			core.FailResp(c, cd.InvalidParams)
			return
		}

		claims, err := utils.ParseToken(Authorization)
		if err != nil {
			core.FailResp(c, cd.ErrorAuthToken)
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			core.FailResp(c, cd.TimeoutAuthToken)
			return
		}

		c.Next()
	}

}
