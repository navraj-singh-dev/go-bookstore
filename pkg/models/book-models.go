package models

import (
	gorm "github.com/jinzhu/gorm"
	config "github.com/navraj-singh-dev/go-bookstore/pkg/config"
)

var db *gorm.DB

type BookModel struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&BookModel{})
}

// --------- model functions ----------
// func (b *BookModel) CreateBook() *BookModel {
// 	db.NewRecord(b)
// 	db.Create(&b)
// 	return b
// }

func CreateBook(newBook *BookModel) *BookModel {
	db.NewRecord(newBook)
	db.Create(newBook)
	return newBook
}

func GetAllBooks() []BookModel {
	var AllBooks []BookModel
	db.Find(&AllBooks)
	return AllBooks
}

func GetBookById(id int64) (*BookModel, *gorm.DB) {
	var RetrievedBook *BookModel
	db := db.Where("ID=?", id).Find((&RetrievedBook))
	return RetrievedBook, db
}

func DeleteById(id int64) BookModel {
	var DeletedBook BookModel
	db.Where("ID=?", id).Delete(&DeletedBook)
	return DeletedBook
}
