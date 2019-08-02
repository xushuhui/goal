package core

const (
	OK              = iota
	DB_CONNECT_ERROR 
	SQL_ERROR        
	INVALID_PARAMS   
	SYSTEM_ERROR     
	NO_USER         
	ERROR_PASSWORD  
	USER_EXIST     
)

var errormap map[int]string = map[int]string{
	OK:               "OK",
	DB_CONNECT_ERROR: "db connect error",
	SQL_ERROR:        "db sql error",
	INVALID_PARAMS:   "error params",
	SYSTEM_ERROR:     "system error",
	
	NO_USER:          "no user",
	ERROR_PASSWORD:   "pass error",
	USER_EXIST:       "user exist",
}

func GetMsg(errorcode int) string {
	return errormap[errorcode]
}
