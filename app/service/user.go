package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-web/app/model"
	"go-web/app/request"
	"go-web/core"
	"net/http"
)

var userModel *model.User

var response *core.Response

func UserLogin(writer http.ResponseWriter, userRequest *request.UserLoginStruct) {
	model.GetUserByAcc(userRequest.Mobile)
	fmt.Println(userModel.UserID)
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
	response.SetData(writer, data)
	return
}
func UserReg(writer http.ResponseWriter, userRequest *request.UserRegStruct) {

	//if result := userModel.GetUserByAcc(userRequest.Mobile); result != false {
	//	response.Fail(writer, core.USER_EXIST)
	//	return
	//}

}
