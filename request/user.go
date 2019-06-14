package request

import (
	"go-web/core"
	"net/http"
)

type UserLoginStruct struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserLogin(Request *http.Request) *UserLoginStruct {
	var user *UserLoginStruct
	err := core.Bind(Request, &user)
	if err != nil {
		panic(err)
	}
	return user
}

type UserRegStruct struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserReg(Request *http.Request) *UserRegStruct {
	var user *UserRegStruct
	err := core.Bind(Request, &user)
	if err != nil {
		panic(err)
	}
	return user
}
