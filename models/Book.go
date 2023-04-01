package models

import (
	"github.com/Bonifaceebuka/book-store-RESTful-API-in-GO/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Id          uint   `gorm:"primary_key;autoIncrement" json: "id"`
	Name        string `json: "name"`
	Author      string `json: "author"`
	Publication string `json: "publication"`
}

func init() {
	db = config.GetDBConnection()
}

func GetAllBooks() []Book {
	var books []Book
	db.Model(&books).Find(&books)
	return books
}

func GetAllBookById(id int) (*Book, *gorm.DB) {
	var book Book
	db.Where("id=?", id).Find(&book)
	return &book, db
}
