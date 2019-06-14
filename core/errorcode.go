package core

const (
	CODE_DB_CONNECT_ERROR int = -4
	CODE_SQL_ERROR        int = -3
	CODE_INVALID_PARAMS   int = -2
	CODE_SYSTEM_ERROR     int = -1
	CODE_OK               int = 0
	CODE_NO_USER          int = 1000
	CODE_ERROR_PASSWORD   int = 1001
	CODE_USER_EXIST       int = 1002
)

var errormap map[int]string = map[int]string{
	CODE_DB_CONNECT_ERROR: "db connect error",
	CODE_SQL_ERROR:        "db sql error",
	CODE_INVALID_PARAMS:   "error params",
	CODE_SYSTEM_ERROR:     "system error",
	CODE_OK:               "OK",
	CODE_NO_USER:          "no user",
	CODE_ERROR_PASSWORD:   "pass error",
	CODE_USER_EXIST:       "user exist",
}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}
