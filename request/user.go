package request
import (
    "fmt"
    "encoding/json"
    "net/http"
)
type User struct {
    Mobile string `json:"mobile"`
    Pwd string `json:"password"`
}
func UserLogin(Request *http.Request) *User {
    var user *User
    err := json.NewDecoder(Request.Body).Decode(&user)
    if err != nil {
        fmt.Println(err);
    }
    return user
}