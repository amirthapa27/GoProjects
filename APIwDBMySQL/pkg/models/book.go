package models

import (
	"mysqlAPI/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() { //intializing the DB
	config.Connect()        //help us connect db
	db = config.GetDB()     //return db we get from the config file
	db.AutoMigrate(&Book{}) //create table
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBook(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	// db.First(&getBook, "ID=?", id)
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var getBook Book
	db.Where("ID = ?", Id).Delete(&getBook)
	return getBook
}

// func UpdateBook(Id int64) (*Book, *gorm.DB) {
// var getBook Book
// db := db.Where("ID = ?", Id).Update(&getBook)
// return &getBook, db
// }
