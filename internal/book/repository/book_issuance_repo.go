package repository

import (
	"lms/database"
	"lms/models"
)

// BookIssuanceRepository defines repository methods
type BookIssuanceRepository interface {
	CreateBookIssuance(issuance *models.BookIssuance) error
	GetAllIssuedBooks() ([]models.BookIssuance, error)
	GetBookIssuanceByID(id string) (models.BookIssuance, error)
	UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error)
	DeleteBookIssuance(id string) error
}

// bookIssuanceRepo implements BookIssuanceRepository
type bookIssuanceRepo struct{}

func NewBookIssuanceRepo() BookIssuanceRepository {
	return &bookIssuanceRepo{}
}

func (r *bookIssuanceRepo) CreateBookIssuance(issuance *models.BookIssuance) error {
	return database.DB.Create(issuance).Error
}

func (r *bookIssuanceRepo) GetAllIssuedBooks() ([]models.BookIssuance, error) {
	var books []models.BookIssuance
	result := database.DB.Find(&books)
	return books, result.Error
}

func (r *bookIssuanceRepo) GetBookIssuanceByID(id string) (models.BookIssuance, error) {
	var book models.BookIssuance
	result := database.DB.First(&book, id)
	return book, result.Error
}

func (r *bookIssuanceRepo) UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error) {
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

func (r *bookIssuanceRepo) DeleteBookIssuance(id string) error {
	return database.DB.Delete(&models.BookIssuance{}, id).Error
}
