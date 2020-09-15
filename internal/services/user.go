package services

import (
	"goal/internal/model"
	"goal/internal/request"
	"goal/pkg/core"
	"goal/pkg/errcode"
	"goal/pkg/lib"
	"golang.org/x/crypto/bcrypt"
)

func Login(req request.Login) (data string, err error) {

	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	if err != nil {
		return
	}
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		err = core.NewError(errcode.ErrorPassWord)
		return
	}
	data, err = lib.GenerateToken(userModel.ID)
	if err != nil {
		return
	}

	return
}
