package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-web/core"
	"go-web/model"
	"go-web/request"
	"net/http"
	//	"fmt"
)

func UserLogin(writer http.ResponseWriter, userRequest *request.User) {
	db, err := gorm.Open("mysql", "root:root@/zhj?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		core.Fail(writer, core.CODE_DB_CONNECT_ERROR)
		return
	}
	db.SingularTable(true)

	var userModel []model.User
	acc := db.Where("user_acc = ?", userRequest.Mobile).First(&userModel)

	if acc.Error != nil {
		core.FailMsg(writer, core.CODE_DB_CONNECT_ERROR, acc.Error.Error())
		return
	}
	if len(userModel) == 0 {
		core.Fail(writer, core.CODE_NO_USER)
		return
	}

	userInfo := userModel[0]
	if core.ValidatePassword(userRequest.Pwd, userInfo.UserEntry, userInfo.UserPwd) == false {
		core.Fail(writer, core.CODE_ERROR_PASSWORD)
		return
	}
	data := make(map[string]interface{})

	data["id"] = 1
	data["token"] = "test"
	core.SetData(writer, data)
	defer db.Close()
	return
}
