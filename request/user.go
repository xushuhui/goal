package request

import (
	"go-web/core"
	"net/http"
)

type User struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

func UserLogin(Request *http.Request) *User {
	var user *User
	err := core.Bind(Request, &user)
	if err != nil {
		panic(err)
	}
	return user
}
