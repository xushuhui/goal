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

	_ = routes.Router().Run(":8000")

}
