package models

import (
	"github.com/Bonifaceebuka/jwt-authentication-backend-service/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password []byte `json: "_"`
	Image    string `json: "image"`
}

func init() {
	db = config.GetDBConnection()
	db.AutoMigrate(&User{})
}

func GetAllUsers() []User {
	var users []User
	db.Model(&users).Find(&users)
	return users
}

func GetAllUserById(id int) (*User, *gorm.DB) {
	var user User
	db.Where("id=?", id).Find(&user)
	return &user, db
}
