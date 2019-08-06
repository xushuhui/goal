package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-web/app/request"
	"go-web/core"
)

//var userModel *models.User

func UserLogin(loginRequest *request.UserLoginRequest) []byte {
	//models.GetUserByAcc(userRequest.Mobile)
	//fmt.Println(userModel.UserID)
	//if result:= userModel.GetUserByAcc(userRequest.Mobile); result == false {
	//	response.Fail(writer, core.NO_USER)
	//	return
	//}
	//
	//if !core.ValidatePassword(userRequest.Pwd, userModel.UserEntry, userModel.UserPwd) {
	//	response.Fail(writer, core.ERROR_PASSWORD)
	//	return
	//}
	data := make(map[string]interface{})

	data["id"] = 1
	data["token"] = "test"
	fmt.Println(data)
	return core.SetData(data)
}
func UserReg(regRequest request.UserRegRequest) {

	//if result := userModel.GetUserByAcc(userRequest.Mobile); result != false {
	//	response.Fail(writer, core.USER_EXIST)
	//	return
	//}
	core.Succeed()
	return
}
