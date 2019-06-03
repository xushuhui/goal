package service
import (
	"net/http"	
	"go-web/core"
	"go-web/request"
	"fmt"
)
func UserLogin(writer http.ResponseWriter,user *request.User)  {
	fmt.Println(user)
	if(user.Mobile != "13012341234"){
		core.Fail(writer,core.CODE_NO_USER)
	}
	if(user.Pwd != "123456"){
		core.Fail(writer,core.CODE_ERROR_PASSWORD)
	}
	core.Fail(writer,core.CODE_SYSTEM_ERROR)
}