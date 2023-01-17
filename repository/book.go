package repository

import (
	"book-api/db"
	model "book-api/model"
)

func AddBook(book model.Book) error {
	if err := db.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBooks() ([]model.Book, error) {
	results := []model.Book{}
	if err := db.DB.Find(&results).Error; err != nil {
		return []model.Book{}, err
	}
	return results, nil
}

func GetBook(id int) (model.Book, error) {
	result := model.Book{}
	if err := db.DB.Where("id  = ?", id).Find(&result).Error; err != nil {
		return model.Book{}, err
	}
	return result, nil
}

func UpdateBook(id int, book model.Book) error {
	if err := db.DB.Model(&model.Book{}).Where("id = ?", id).Updates(model.Book{Title: book.Title, Author: book.Author, Category: book.Category}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBook(id int) error {
	if err := db.DB.Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
