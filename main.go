package main

import (
	"github.com/jinzhu/gorm"
	"go-web/core"
	"go-web/routes"
)

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "zhj_" + defaultTableName
	}

}
func main() {
	core.InitEnv()
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// })

	//test(w http.ResponseWriter)

	// http.HandleFunc("/", httpHandle)

	//http.HandleFunc("/user/login", controller.UserLogin)
	//http.HandleFunc("/user/reg", controller.UserReg)
	_ = routes.Router().Run(":8000")
	//core.ReadMysqlConf()
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	log.Fatal("Listen", err)
	//}
}
