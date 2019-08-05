package core

const (
	OK = iota
	DBCONNECTERROR
	SQLERROR
	INVALIDPARAMS
	SYSTEMERROR
	NOUSER
	ERRORPASSWORD
	USEREXIST
)

var errormap map[int]string = map[int]string{
	OK:             "OK",
	DBCONNECTERROR: "db connect error",
	SQLERROR:       "db sql error",
	INVALIDPARAMS:  "error params",
	SYSTEMERROR:    "system error",
	NOUSER:         "no user",
	ERRORPASSWORD:  "pass error",
	USEREXIST:      "user exist",
}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}
