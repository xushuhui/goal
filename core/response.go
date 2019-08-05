package core

import (
	"encoding/json"
	"log"
)

type Response struct {
	ErrorCode int         `json:"error_code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
}
type Data struct {
	Data interface{} `json:"data,omitempty"`
}

func Fail(error_code int) []byte {
	return Resp(error_code, GetMsg(error_code), nil)
}
func FailMsg(error_code int, msg string) []byte {
	return Resp(error_code, msg, nil)
}
func Succeed() []byte {
	return Resp(OK, GetMsg(OK), nil)
}
func SetData(data interface{}) []byte {
	return Resp(OK, GetMsg(OK), data)
}

func Resp(error_code int, msg string, data interface{}) []byte {
	//writer.Header().Set("Content-Type", "application/json")
	//writer.WriteHeader(http.StatusOK) //设置200状态
	res := Response{
		ErrorCode: error_code,
		Msg:       msg,
		Data:      data,
	}
	ret, err := json.Marshal(res)
	if err != nil {
		log.Println(err.Error())
	}

	return ret
}
