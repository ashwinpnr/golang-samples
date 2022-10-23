package models

import (
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func (b *Book) ToString() string {
	bookStr := "Name : " + b.Name + " Author : " + b.Author + " Publication : " + b.Publication
	return bookStr
}

func init() {
	config.ConnectDB()
	DB = config.GetDBConn()
	DB.AutoMigrate(&Book{})
}

func CreateBook(b Book) *Book {
	DB.Create(&b)
	return &b

}

func GetAllBooks() []Book {
	var books []Book
	DB.Find(&books)
	return books
}

func GetBookByID(bookId int64) Book {
	var book Book
	DB.First(&book, bookId)
	return book
}

func DeleteBook(bookId int64) Book {
	var book Book
	DB.Delete(&book, bookId)
	return book
}

func UpdateBook(bookId int64, b Book) *Book {
	existingBook := GetBookByID(bookId)
	if b.Name != "" {
		existingBook.Name = b.Name
	}
	if b.Author != "" {
		existingBook.Author = b.Author
	}
	if b.Publication != "" {
		existingBook.Publication = b.Publication
	}
	DB.Save(&existingBook)
	return &existingBook
}
