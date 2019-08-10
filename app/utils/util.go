/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func BindJson(c *gin.Context, obj interface{}) error {
	s, err := ioutil.ReadAll(c.Request.Body) //把  body 内容读入字符串
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
