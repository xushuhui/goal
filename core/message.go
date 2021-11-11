package core

const (
	SUCCESS = iota
	InvalidParams
	DBError
	ServerError
)

var MsgFlags = map[int]string{
	SUCCESS:       "ok",
	InvalidParams: "请求参数错误",
	DBError:       "数据库错误",
	ServerError:   "系统异常，请联系管理员！",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[InvalidParams]
}
