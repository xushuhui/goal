package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goal/global"
	"goal/pkg/core"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		e := c.Errors.Last()
		if e == nil {
			return
		}
		err := e.Err
		var errStr string
		//// 记录错误日志
		global.Logger.Warn(err)
		switch err.(type) {

		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errStr = fmt.Errorf("%s 类型错误，期望类型 %s", unmarshalTypeError.Field, unmarshalTypeError.Type.String()).Error()
		default:
			errStr = e.Err.Error()
		}

		core.InvalidParamsResp(c, errStr)
	}

}
