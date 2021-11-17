package core

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Error 数据返回通用 JSON 数据结构
type IError struct {
	Code    int         `json:"code"`    // 错误码 ((0: 成功，1: 失败，>1: 错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据 (业务接口定义具体数据结构)
	Err     error       `json:"-"`
}

type HttpError struct {
	IError
	Status int
}

func (err IError) Error() (re string) {
	return fmt.Sprintf("code=%v, Message=%v,Err=%v", err.Code, err.Message, err.Err)
}

func NewIError(code int, message string) IError {
	return IError{
		Code:    code,
		Message: message,
	}
}

func NewHttpError(code int) (err error) {
	err = HttpError{
		NewIError(code, GetMsg(code)),
		fiber.StatusInternalServerError,
	}
	return
}

func NewErrorMessage(code int, message string) (err error) {
	err = HttpError{
		NewIError(code, message),
		fiber.StatusBadRequest,
	}
	return
}

func ParamsError(code int) (err error) {
	return HttpError{
		NewIError(code, GetMsg(code)),
		fiber.StatusBadRequest,
	}
}

func NotFoundError(code int) error {
	return HttpError{
		NewIError(code, GetMsg(code)),
		fiber.StatusNotFound,
	}
}

func UnAuthenticatedError(code int) error {
	return HttpError{
		NewIError(code, GetMsg(code)),
		fiber.StatusUnauthorized,
	}
}

func ForbiddenError(code int) error {
	return HttpError{
		NewIError(code, GetMsg(code)),
		fiber.StatusForbidden,
	}
}
