package ext

import (
	"github.com/gin-gonic/gin"

	"strconv"
)

var Context *gin.Context

type Input struct {
	body string
}

func (input *Input) get(p string) *Input {

	d, e := Context.GetQuery(p)
	input.body = d
	if e == false {
		return input
	}
	return input
}

func (input *Input) post(p string) *Input {
	d, e := Context.GetPostForm(p)

	input.body = d
	if e == false {
		return input
	}

	return input
}

func (input *Input) String() string {
	return input.body
}

func (input *Input) Atoi() int {
	body, _ := strconv.Atoi(input.body)
	return body
}

//封装之前
//pid, _ := strconv.Atoi(c.Query("product_id"))
//alias := c.Query("product_alias")
//
//封装之后
//pid := input.Get("product_id").Atoi()
//alias := input.Get("product_alias").String()
