/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package core

const (
	OK = iota
	DBCONNECTERROR
	SQLERROR
	INVALIDPARAMS
	SYSTEMERROR
	NOUSER
	ERRORPASSWORD
	USEREXIST
	TOKENMUST
)

var codeMap = map[int]string{
	OK:             "OK",
	DBCONNECTERROR: "db connect error",
	SQLERROR:       "db sql error",
	INVALIDPARAMS:  "error params",
	SYSTEMERROR:    "system error",
	NOUSER:         "no user",
	ERRORPASSWORD:  "pass error",
	USEREXIST:      "user exist",
	TOKENMUST:      "token must",
}

func GetMsg(code int) string {
	return codeMap[code]
}
