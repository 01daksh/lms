package repository

import (
	"lms/database"
	"lms/models"
)

func CreateBookIssuance(issuance *models.BookIssuance) error {
	result := database.DB.Create(issuance)
	return result.Error
}

func GetAllIssuedBooks() ([]models.BookIssuance, error) {
	var books []models.BookIssuance
	result := database.DB.Find(&books)
	return books, result.Error
}

func GetBookIssuanceByID(id string) (models.BookIssuance, error) {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	return book, result.Error
}

func UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error) {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return book, result.Error
	}

	book.BookTitle = updatedBook.BookTitle
	book.IssuedTo = updatedBook.IssuedTo
	book.IssuedDate = updatedBook.IssuedDate
	book.ReturnDate = updatedBook.ReturnDate

	result = database.DB.Save(&book)
	return book, result.Error
}

func DeleteBookIssuance(id string) error {
	result := database.DB.Delete(&models.BookIssuance{}, id)
	return result.Error
}
