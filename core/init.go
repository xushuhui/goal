package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"os"

	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var redisConn *redis.Client

func InitEnv() {
	GetDBConn()
	initRedis()
}

func initRedis() {
	conf := ReadRedisConf()
	redisConn := redis.NewClient(&redis.Options{
		Addr:     conf.Host,
		Password: "",
		DB:       0,
	})

	_, err := redisConn.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
	}

}

// DBConn : 返回数据库连接对象
func GetDBConn() *gorm.DB {

	conf := ReadMysqlConf()
	mysqlSource := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ")/" + conf.Dbname + "?charset=utf8"

	DbConn, err := gorm.Open("mysql", mysqlSource)

	DbConn.SingularTable(true)
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}

	gorm.DefaultTableNameHandler = func(DbConn *gorm.DB, defaultTableName string) string {
		return "zhj_" + defaultTableName
	}
	return DbConn
}
func RedisConn() *redis.Client {
	return redisConn
}
