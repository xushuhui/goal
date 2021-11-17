package core

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(obj interface{}) error {
	err := validate.Struct(obj)
	if err != nil {
		s := Translate(err.(validator.ValidationErrors))
		return NewErrorMessage(InvalidParams, s)
	}
	return nil
}

func ParseQuery(c *fiber.Ctx, request interface{}) (err error) {
	err = c.QueryParser(request)

	if err != nil {
		return err
	}

	err = ValidateRequest(request)
	return
}

func ParseRequest(c *fiber.Ctx, request interface{}) (err error) {
	err = c.BodyParser(request)

	if err != nil {
		return err
	}
	err = ValidateRequest(request)
	return
}

func SuccessResp(c *fiber.Ctx) error {
	return c.JSON(IError{
		Code:    0,
		Message: GetMsg(0),
	})
}

func SetData(c *fiber.Ctx, data interface{}) error {
	return c.JSON(IError{
		Code:    0,
		Message: GetMsg(0),
		Data:    data,
	})
}

func SetPage(c *fiber.Ctx, list interface{}, totalRows int) error {
	return c.JSON(IError{
		Code:    0,
		Message: GetMsg(0),
		Data: Pager{
			Page:  GetPage(c),
			Size:  GetSize(c),
			Total: totalRows,
			Items: list,
		},
	})
}

func HandleServerError(c *fiber.Ctx, err error) error {
	return c.JSON(IError{
		Code:    ServerError,
		Message: GetMsg(ServerError),
		Err:     err,
	})
}

func (err *HttpError) HandleHttpError(c *fiber.Ctx) error {
	return c.Status(err.Status).JSON(IError{
		Code: err.Code, Message: err.Message,
	})
}
