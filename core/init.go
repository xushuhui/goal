package core

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var dbConn *sql.DB
var redisConn *redis.Client

func InitEnv() {
	initMysql()
	initRedis()
}
func initMysql() {
	conf := ReadMysqlConf()
	//"test:test@tcp(127.0.0.1:3306)/fileserver?charset=utf8"
	mysqlSource := conf.User + ":" + conf.Pwd + "@tcp(" + conf.Host + ")/" + conf.Dbname + "?charset=utf8"
	dbConn, _ = sql.Open("mysql", mysqlSource)
	dbConn.SetMaxOpenConns(1000)
	err := dbConn.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}
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
func DBConn() *sql.DB {
	return dbConn
}
func RedisConn() *redis.Client {
	return redisConn
}
