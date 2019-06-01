package controller
import (
	"net/http"
	
	
	"go-web/core"

)


//UserLogin func
func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	// request.ParseForm()

	// mobile := request.PostForm.Get("mobile")
	// passwd := request.PostForm.Get("passwd")

	//loginok := true
	core.Fail(writer,core.CODE_SYSTEM_ERROR)
	// if(loginok){
	// 	data := make(map[string]interface{})
	// 	data["id"]=1
	// 	data["token"]="test"
	// 	core.Succeed(writer)
	// }

}