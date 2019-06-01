package controller
import (
	"net/http"
	
	
	"answer-server/core"

)
func UserLogin(writer http.ResponseWriter,
	request *http.Request) {
	// request.ParseForm()

	// mobile := request.PostForm.Get("mobile")
	// passwd := request.PostForm.Get("passwd")

	loginok := true
	
	if(loginok){
		//{"id":1,"token":"xx"}
		data := make(map[string]interface{})
		data["id"]=1
		data["token"]="test"
		core.Succeed(writer)
	}

}