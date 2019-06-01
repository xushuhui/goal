package core

const(
	CODE_SQL_ERROR int = -3
	CODE_INVALID_PARAMS int = -2
	CODE_SYSTEM_ERROR int = -1
	CODE_OK int = 0
	CODE_NO_USER int = 1000
)
var errormap map[int]string = map[int]string{
	CODE_OK:"OK",
	CODE_SYSTEM_ERROR:"system error"}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}