package service
import (
	"net/http"	
	"go-web/core"
	"go-web/request"
	"go-web/model"
	"github.com/jinzhu/gorm"
	 _ "github.com/go-sql-driver/mysql"
//	"fmt"

)
func UserLogin(writer http.ResponseWriter,userRequest *request.User)  {
	db, err := gorm.Open("mysql", "root:root@/zhj?charset=utf8&parseTime=True&loc=Local")
	if (err != nil){
		core.Fail(writer,core.CODE_DB_CONNECT_ERROR)
		return 
	} 
	db.SingularTable(true)
	defer db.Close()
	var userModel []model.User
	acc := db.Where("user_acc = ?", userRequest.Mobile).First(&userModel)
	
	if(acc.Error!=nil){
		core.FailMsg(writer,core.CODE_DB_CONNECT_ERROR,acc.Error.Error())
		return 
	}
	if(len(userModel) == 0){
		core.Fail(writer,core.CODE_NO_USER)
		return 
	}
//	fmt.Printf("v1 type:%T\n", userModel)
//	fmt.Println(core.Md5Encode("123456"))
	
	if(core.ValidatePassword(userRequest.Pwd,userModel[0].UserEntry,userModel[0].UserPwd)){
		core.Fail(writer,core.CODE_ERROR_PASSWORD)
		return 
	}
	data := make(map[string]interface{})
	data["id"]=1
	data["token"]="test"
	core.SetData(writer,data)
	return 
}