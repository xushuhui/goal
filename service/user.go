package service
import (
	"net/http"	
	"go-web/core"
	"go-web/request"
)
func UserLogin(writer http.ResponseWriter,user *request.User)  {
	if(user.Mobile != "13012341234"){
		core.Fail(writer,core.CODE_NO_USER)
		return 
	}
	if(user.Pwd != "123456"){
		core.Fail(writer,core.CODE_ERROR_PASSWORD)
		return 
	}
	data := make(map[string]interface{})
	data["id"]=1
	data["token"]="test"
	core.SetData(writer,data)
	return 
}