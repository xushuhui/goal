package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"go-web/app/request"
	"go-web/core"
	"net/http"
)

//var userModel *models.User

func UserLogin(c *gin.Context) {
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
	core.SetData(data)
	return
}
func UserReg(writer http.ResponseWriter, userRequest *request.UserRegStruct) {

	//if result := userModel.GetUserByAcc(userRequest.Mobile); result != false {
	//	response.Fail(writer, core.USER_EXIST)
	//	return
	//}
	core.Succeed()
	return
}
