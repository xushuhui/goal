package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	UserID    uint   `gorm:"primary_key"`
	UserAcc   string `gorm:"column:user_acc"`
	UserPwd   string
	UserEntry string
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
func GetUserByAcc(mobile string) User {
	var User = User{}

	//err :=db.Where("user_acc = ?", mobile).First(&User).Error
	//if err == gorm.ErrRecordNotFound {
	//	return nil
	//}
	return User
}
