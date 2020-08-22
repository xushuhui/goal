package core

import (
	"goal/global"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var RDB *redis.Client

// 初始化连接
func initRedis() (err error) {
	NewClient()
	_, err := RDB.Ping().Result()
	if err != nil {
		return
	}

}
func NewClient() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password, // no password set
		DB:       0,                            // use default DB
	})
}
func SelectDB(db int) {
	RDB.Do("SELECT", db)
}
