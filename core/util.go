package core

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func Bind(req *http.Request, obj interface{}) error {
	contentType := req.Header.Get("Content-Type")
	//如果是简单的json
	if strings.Contains(strings.ToLower(contentType), "application/json") {
		return BindJson(req, obj)
	}
	return errors.New("当前方法暂不支持")
}

func BindJson(req *http.Request, obj interface{}) error {
	s, err := ioutil.ReadAll(req.Body) //把  body 内容读入字符串
	if err != nil {
		return err
	}
	err = json.Unmarshal(s, obj)
	return err
}
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func MakePassword(plainpwd, salt string) string {
	return Md5Encode(Md5Encode(plainpwd) + Md5Encode(salt))
}
func ValidatePassword(plainpwd, salt, password string) bool {
	return MakePassword(plainpwd, salt) == password
}
