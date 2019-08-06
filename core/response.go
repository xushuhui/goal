package core

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Fail(code int) *Response {
	return Resp(code, GetMsg(code), nil)
}
func FailMsg(code int, msg string) *Response {
	return Resp(code, msg, nil)
}
func Succeed() *Response {
	return Resp(OK, GetMsg(OK), nil)
}
func SetData(data interface{}) *Response {
	return Resp(OK, GetMsg(OK), data)
}

func Resp(code int, msg string, data interface{}) *Response {

	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
func JsonResp(c *gin.Context, resp *Response) {
	c.JSON(http.StatusOK, resp)
}
