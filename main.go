package main

import (
	"go-web/controller"
	"net/http"
	//"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"log"
)

func init() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "zhj_" + defaultTableName
	}

}
func main() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// })

	//test(w http.ResponseWriter)
	// conf := config.ReadConfig();
	// log.Println("conf", conf)
	// http.HandleFunc("/", httpHandle)
	http.HandleFunc("/user/login", controller.UserLogin)
	http.HandleFunc("/user/reg", controller.UserReg)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Listen", err)
	}
}
