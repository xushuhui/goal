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

func Fail(writer http.ResponseWriter, error_code int) {
	Resp(writer, error_code, GetMsg(error_code), nil)
}
func FailMsg(writer http.ResponseWriter, error_code int,msg string) {
	Resp(writer, error_code, msg, nil)
}
func Succeed(writer http.ResponseWriter) {
	Resp(writer, CODE_OK, GetMsg(CODE_OK), nil)
}
func SetData(writer http.ResponseWriter, data interface{}) {
	Resp(writer, CODE_OK, GetMsg(CODE_OK), data)
}

func Resp(writer http.ResponseWriter, error_code int, msg string, data interface{}) {
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
