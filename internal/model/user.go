package model

import (
	"goal/global"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Phone    string `gorm:"column:phone"`
	Password string `gorm:"column:password"`
}

func GetAccountUserOne(where string, args ...interface{}) (model User, err error) {
	err = global.DBEngine.First(&model, where, args).Error
	return

}
