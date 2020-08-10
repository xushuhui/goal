package services

import (
	"github.com/gin-gonic/gin"
	"goal/app/model"
	"goal/app/request"
	"goal/core"
	"goal/lib"
	errcode2 "goal/pkg/errcode"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context, req request.Login) (err error) {

	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	if err != nil {
		return
	}
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(c, errcode2.ErrorPassWord)
		return
	}
	data, err := lib.GenerateToken(userModel.Id)
	if err != nil {
		return
	}
	core.SetData(c, data)
	return
}
