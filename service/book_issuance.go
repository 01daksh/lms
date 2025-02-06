package service

import (
	"lms/models"
	"lms/repository"
)

func CreateBookIssuance(issuance *models.BookIssuance) error {
	return repository.CreateBookIssuance(issuance)
}

func GetAllIssuedBooks() ([]models.BookIssuance, error) {
	return repository.GetAllIssuedBooks()
}

func GetBookIssuanceByID(id string) (models.BookIssuance, error) {
	return repository.GetBookIssuanceByID(id)
}

func UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error) {
	return repository.UpdateBookIssuance(id, updatedBook)
}

func DeleteBookIssuance(id string) error {
	return repository.DeleteBookIssuance(id)
}
