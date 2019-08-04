package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID    uint   `gorm:"primary_key"`
	UserAcc   string `gorm:"column:user_acc"`
	UserPwd   string `gorm:"column:user_pwd"`
	UserEntry string `gorm:"column:user_entry"`
}

func (User) TableName() string {
	return "user"
}

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:root@/zhj?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)
}
func GetUserByAcc(mobile string) (res interface{}) {
	var user User
	err := db.Where("user_acc = ?", mobile).First(&user)
	return err

}
