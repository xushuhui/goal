package controller
import (
	"net/http"
	 "go-web/request"
	"go-web/service"
)


//UserLogin func
func UserLogin(writer http.ResponseWriter,r *http.Request) {
	//  request.ParseForm()
	// mobile :=request.PostForm.Get("mobile")
	// pwd:=request.PostForm.Get("password")
	
    if r.Body == nil {
        http.Error(writer, "Please send a request body", 400)
        return
	}
	user := request.UserLogin(r);
	service.UserLogin(writer,user)
	
	//core.Fail(writer,core.CODE_SYSTEM_ERROR)
	

}