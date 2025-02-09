package service

import (
	"lms/internal/book/repository"
	"lms/models"
)

// BookIssuanceService defines service methods
type BookIssuanceService interface {
	CreateBookIssuance(issuance *models.BookIssuance) error
	GetAllIssuedBooks() ([]models.BookIssuance, error)
	GetBookIssuanceByID(id string) (models.BookIssuance, error)
	UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error)
	DeleteBookIssuance(id string) error
}

// bookIssuanceService implements BookIssuanceService
type bookIssuanceService struct {
	repo repository.BookIssuanceRepository
}

func NewBookIssuanceService(repo repository.BookIssuanceRepository) BookIssuanceService {
	return &bookIssuanceService{repo: repo}
}

func (s *bookIssuanceService) CreateBookIssuance(issuance *models.BookIssuance) error {
	return s.repo.CreateBookIssuance(issuance)
}

func (s *bookIssuanceService) GetAllIssuedBooks() ([]models.BookIssuance, error) {
	return s.repo.GetAllIssuedBooks()
}

func (s *bookIssuanceService) GetBookIssuanceByID(id string) (models.BookIssuance, error) {
	return s.repo.GetBookIssuanceByID(id)
}

func (s *bookIssuanceService) UpdateBookIssuance(id string, updatedBook *models.BookIssuance) (models.BookIssuance, error) {
	return s.repo.UpdateBookIssuance(id, updatedBook)
}

func (s *bookIssuanceService) DeleteBookIssuance(id string) error {
	return s.repo.DeleteBookIssuance(id)
}
