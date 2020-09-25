package services

import (
	"goal/internal/model"
	"goal/internal/request"
	"goal/pkg/core"
	"goal/pkg/errcode"
	"goal/pkg/lib"
	"golang.org/x/crypto/bcrypt"
)

func Login(req request.Login) (data map[string]interface{}, e error) {

	userModel, e := model.GetAccountUserOne("phone=?", req.Phone)
	if e != nil {
		return
	}
	// 正确密码验证
	e = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if e != nil {
		e = core.NewError(errcode.ErrorPassWord)
		return
	}
	token, e := lib.GenerateToken(userModel.ID)
	if e != nil {
		return
	}
	data = make(map[string]interface{})
	data["access_token"] = token
	return
}
