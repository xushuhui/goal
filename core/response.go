package core

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
}
type Data struct {
	Data interface{} `json:"data,omitempty"`
}

func (response *Response)Fail(writer http.ResponseWriter, error_code int) {
	response.Resp(writer, error_code, GetMsg(error_code), nil)
}
func (response *Response)FailMsg(writer http.ResponseWriter, error_code int, msg string) {
	response.Resp(writer, error_code, msg, nil)
}
func (response *Response)Succeed(writer http.ResponseWriter) {
	response.Resp(writer, OK, GetMsg(OK), nil)
}
func (response *Response)SetData(writer http.ResponseWriter, data interface{}) {
	response.Resp(writer, OK, GetMsg(OK), data)
}

func (response *Response)Resp(writer http.ResponseWriter, error_code int, msg string, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK) //设置200状态
	res := Response{
		ErrorCode: error_code,
		Msg:       msg,
		Data:      data,
	}
	ret, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}
	writer.Write(ret)
	return
}
