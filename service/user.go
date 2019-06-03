package service
import (
	"net/http"	
	"go-web/core"
	"go-web/request"
	"fmt"
)
func UserLogin(user *request.User,writer http.ResponseWriter)  {
	fmt.Println(user)
	core.Fail(writer,core.CODE_SYSTEM_ERROR)
}