package models

import (
	"github.com/KleitonBarone/Go-Bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `gorm:""json:"author"`
	Publication string `gorm:""json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (newBook *Book) CreateBook() *Book {
	db.NewRecord(newBook)
	db.Create(&newBook)
	return newBook
}

func GetAllBooks() []Book {
	var allBooks []Book
	db.Find(&allBooks)
	return allBooks
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var deletedBook Book
	db.Where("ID=?", Id).Delete(deletedBook)
	return deletedBook
}
