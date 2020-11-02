package services

import (
	"goal/internal/request"
	"goal/pkg/core"
	"goal/pkg/msg"
)

func Login(req request.Login) (data map[string]interface{}, e error) {

	core.SendMessage(1, msg.User{
		Name: req.Phone,
	})
	//userModel, e := model.GetAccountUserOne("phone=?", req.Phone)
	//if e != nil {
	//	return
	//}
	//// 正确密码验证
	//e = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password))
	//if e != nil {
	//	e = core.NewError(errcode.ErrorPassWord)
	//	return
	//}
	//token, e := lib.GenerateToken(userModel.ID)
	//if e != nil {
	//	return
	//}
	//data = make(map[string]interface{})
	//data["access_token"] = token
	return
}
