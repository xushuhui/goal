package services

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-web/app/models"
	"go-web/app/request"
	"go-web/core"
	"go-web/core/ext"
)

func UserLogin(loginRequest *request.UserLoginRequest) *core.Response {
	model, err := models.GetUserByAcc(loginRequest.Mobile)
	if err == gorm.ErrRecordNotFound {
		return core.Fail(core.NOUSER)
	}

	if !core.ValidatePassword(loginRequest.Pwd, model.UserEntry, model.UserPwd) {
		return core.Fail(core.ERRORPASSWORD)
	}
	data := make(map[string]interface{})
	data["token"] = ext.GenerateToken(model.UserID)
	return core.SetData(data)
}
func UserReg(regRequest request.UserRegRequest) {

	//if result := userModel.GetUserByAcc(userRequest.Mobile); result != false {
	//	response.Fail(writer, core.USER_EXIST)
	//	return
	//}
	//core.Succeed()
	return
}
