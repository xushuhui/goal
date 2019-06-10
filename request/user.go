package request
import (
    "net/http"
    "go-web/core"
)
type User struct {
    Mobile string `json:"mobile"`
    Pwd string `json:"password"`
}
func UserLogin(Request *http.Request) *User {
    var user *User
    core.Bind(Request,&user)
    return user
}