package core

import (
	"github.com/gin-gonic/gin"
	"goal/pkg/errcode"
)

// JsonResponse 数据返回通用 JSON 数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
}

func ParseRequest(c *gin.Context, request interface{}) (err error) {
	return c.ShouldBind(request)
}
func FailResp(c *gin.Context, code int) {
	c.AbortWithStatusJSON(200, JsonResponse{
		Code:    code,
		Message: errcode.GetMsg(code),
	})
	return
}
func InvalidParamsResp(c *gin.Context, msg string) {

	c.AbortWithStatusJSON(200, JsonResponse{
		Code:    errcode.InvalidParams,
		Message: msg,
	})
	return
}

func SuccessResp(c *gin.Context) {
	c.JSON(200, JsonResponse{
		Code:    0,
		Message: errcode.GetMsg(0),
	})
}
func SetData(c *gin.Context, data interface{}) {
	c.JSON(200, JsonResponse{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data:    data,
	})
}

func SetPage(c *gin.Context, list interface{}, totalRows int) {
	c.JSON(200, JsonResponse{
		Code:    0,
		Message: errcode.GetMsg(0),
		Data: Pager{
			Page:      GetPage(c),
			PageSize:  GetPageSize(c),
			TotalRows: totalRows,
			List:      list,
		},
	})
}
func ServerError(c *gin.Context) {
	c.JSON(500, JsonResponse{
		Code:    errcode.ServerError,
		Message: errcode.GetMsg(errcode.ServerError),
	})
}