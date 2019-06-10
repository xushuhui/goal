package core

import (
	"encoding/json"
	"strings"
	"io/ioutil"
	"errors"
	"net/http"
)
func Bind(req *http.Request,obj interface{}) error{
	contentType := req.Header.Get("Content-Type")
	//如果是简单的json
	if strings.Contains(strings.ToLower(contentType),"application/json"){
		return  BindJson(req,obj)
	}
	return errors.New("当前方法暂不支持")
}

func BindJson(req *http.Request,obj interface{}) error{
	s, err := ioutil.ReadAll(req.Body) //把  body 内容读入字符串
	if err!=nil{
		return err
	}
	err = json.Unmarshal(s,obj)
	return err
}
