package store

import (
	"errors"
	"github.com/JosTENill/go-book-api/models"
)

var books = []models.Book{}
var nextID = 1


func CreateBook(book models.Book) models.Book {
	book.ID = nextID
	nextID++
	books = append(books, book)
	return book
}


func GetBook(id int) (models.Book, error) {
	for _, b := range books {
		if b.ID == id {
			return b, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}


func ListBooks() []models.Book {
	return books
}


func UpdateBook(id int, updated models.Book) (models.Book, error) {
	for i, b := range books {
		if b.ID == id {
			updated.ID = id
			books[i] = updated
			return updated, nil
		}
	}
	return models.Book{}, errors.New("book not found")
}


func DeleteBook(id int) error {
	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}


