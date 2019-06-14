package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"go-web/core"
	"go-web/model"
	"go-web/request"
	"net/http"
)

func UserLogin(writer http.ResponseWriter, userRequest *request.UserLoginStruct) {
	//var userModel *model.User

	userInfo := model.GetUserByAcc(userRequest.Mobile)
	fmt.Println(userInfo)
	fmt.Printf("v1 type:%T\n", userInfo)

	//if userInfo == nil {
	//	core.Fail(writer, core.CODE_NO_USER)
	//	return
	//}

	//if !core.ValidatePassword(userRequest.Pwd, userInfo.UserEntry, userInfo.UserPwd) {
	//	core.Fail(writer, core.CODE_ERROR_PASSWORD)
	//	return
	//}
	data := make(map[string]interface{})

	data["id"] = 1
	data["token"] = "test"
	core.SetData(writer, data)
	return
}
func UserReg(writer http.ResponseWriter, userRequest *request.UserRegStruct) {
	//var userModel []model.User
	//acc := db.Where("user_acc = ?", userRequest.Mobile).First(&userModel)
	userInfo := model.GetUserByAcc(userRequest.Mobile)
	fmt.Println(userInfo)
	fmt.Printf("v1 type:%T\n", userInfo)
	//if userInfo != nil {
	//	core.Fail(writer, core.CODE_USER_EXIST)
	//	return
	//}

}
