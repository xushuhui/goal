package api

import (
	"github.com/gin-gonic/gin"
	"goal/app/code"
	"goal/app/model"
	"goal/app/request"
	"goal/core"
	"goal/lib"
	"golang.org/x/crypto/bcrypt"
)

//UserLogin func
func Login(c *gin.Context) {
	var req request.Login
	if err := core.ParseRequest(c, &req); err != nil {
		c.Error(err)
		return
	}
	userModel, err := model.GetAccountUserOne("phone=?", req.Phone)
	if err != nil {
		c.Error(err)
		return
	}
	// 正确密码验证
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	if err != nil {
		core.FailResp(c, code.ErrorPassWord)
	}
	data, err := lib.GenerateToken(userModel.Id)
	if err != nil {
		c.Error(err)
		return
	}
	core.SetData(c, data)
	return

}
