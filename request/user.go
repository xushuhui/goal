package request

type UserLoginRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}

type UserRegRequest struct {
	Mobile string `json:"mobile"`
	Pwd    string `json:"password"`
}
