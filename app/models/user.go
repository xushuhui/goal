package models

import (
	"go-web/core"
)

type User struct {
	UserID    uint `gorm:"primary_key"`
	UserAcc   string
	UserPwd   string
	UserEntry string
}

func init() {
	//db, err = gorm.Open("mysql", "root:root@/zhj?charset=utf8&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
	//	return "prefix_" + defaultTableName;
	//}
	//db.SingularTable(true)
}

func GetUserByAcc(mobile string) (*User, error) {
	var user User
	err := core.GetDBConn().Where("user_acc = ?", mobile).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUserById(userId int) (*User, error) {
	var user User
	err := core.GetDBConn().First(&user, userId).Error

	if err != nil {
		return nil, err
	}
	return &user, nil
}
