package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goal/config"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(config.DBType, config.DBSource)
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
