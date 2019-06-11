package model
type User struct {
	UserID        uint `gorm:"primary_key"`
	UserAcc string  
	UserPwd string 
	UserEntry string
  }