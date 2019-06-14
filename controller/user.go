package controller

import (
	"go-web/request"
	"go-web/service"
	"net/http"
)

//UserLogin func
func UserLogin(writer http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(writer, "Please send a request body", 400)
		return
	}
	user := request.UserLogin(r)
	service.UserLogin(writer, user)
}
func UserReg(writer http.ResponseWriter, r *http.Request) {
	userRequest := request.UserReg(r)
	service.UserReg(writer, userRequest)
}
