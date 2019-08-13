package request

import (
	"github.com/gin-gonic/gin"
	"go-web/app/utils"
)

type CreateBookRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func CreateBook(c *gin.Context) *CreateBookRequest {
	var request *CreateBookRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}

type UpdateBookRequest struct {
	ID   string `json:"mobile"`
	Name string `json:"password"`
}

func UpdateBook(c *gin.Context) *UpdateBookRequest {
	var request *UpdateBookRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}

type ReadBookRequest struct {
	Id int64 `json:"id"`
}

func ReadBook(c *gin.Context) *ReadBookRequest {
	var request *ReadBookRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}

type DeleteBookRequest struct {
	Id int64 `json:"id"`
}

func DeleteBook(c *gin.Context) *DeleteBookRequest {
	var request *DeleteBookRequest
	err := utils.BindJson(c, &request)
	if err != nil {
		panic(err)
	}
	return request
}
