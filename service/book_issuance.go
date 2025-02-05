package service

import (
	"errors"
	"lms/database"
	"lms/models"

	"gorm.io/gorm"
)

func CreateBookIssuance(issuance *models.BookIssuance) error {
	result := database.DB.Create(issuance)
	return result.Error
}

// this functions return 2 values (list of books issued and error (nil here onn success))
func GetAllIssuedBooks() ([]models.BookIssuance, error) {
	var books []models.BookIssuance
	result := database.DB.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func GetBookIssuanceByID(id string) (models.BookIssuance, error) {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return book, errors.New("record not found")
		}
		return book, result.Error
	}
	return book, nil
}

func UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error) {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return book, errors.New("record not found")
	}

	result = database.DB.Save(updatedBook)
	if result.Error != nil {
		return book, result.Error
	}

	return *updatedBook, nil
}

func DeleteBookIssuance(id string) error {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return errors.New("record not found")
	}

	result = database.DB.Delete(&book)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// we are returning nil so that verification could be done router.go if success or not
