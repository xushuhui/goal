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
	TOKENMUST
)

var codeMap = map[int]string{
	OK:             "OK",
	DBCONNECTERROR: "db connect error",
	SQLERROR:       "db sql error",
	INVALIDPARAMS:  "error params",
	SYSTEMERROR:    "system error",
	NOUSER:         "no user",
	ERRORPASSWORD:  "pass error",
	USEREXIST:      "user exist",
	TOKENMUST:      "token must",
}

func GetMsg(code int) string {
	return codeMap[code]
}
