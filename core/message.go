package core

const (
	SUCCESS = iota
	ServerError
	DBError

	InvalidParams
)

var CodeMapping = map[int]string{
	SUCCESS:     "ok",
	ServerError: "系统异常，请联系管理员！",
	DBError:     "数据库错误",

	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := CodeMapping[code]
	if ok {
		return msg
	}

	return CodeMapping[InvalidParams]
}
