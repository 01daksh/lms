package handlers

import (
	"lms/internal/book/service"
	"lms/models"

	"github.com/gofiber/fiber/v2"
)

// BookIssuanceHandler defines handler methods
type BookIssuanceHandler interface {
	CreateBookIssuance(c *fiber.Ctx) error
	GetAllIssuedBooks(c *fiber.Ctx) error
	GetBookIssuanceByID(c *fiber.Ctx) error
	UpdateBookIssuance(c *fiber.Ctx) error
	DeleteBookIssuance(c *fiber.Ctx) error
}

// bookIssuanceHandler implements BookIssuanceHandler
type bookIssuanceHandler struct {
	service service.BookIssuanceService
}

// NewBookIssuanceHandler is a constructor function
func NewBookIssuanceHandler(service service.BookIssuanceService) BookIssuanceHandler {
	return &bookIssuanceHandler{service: service}
}

// CreateBookIssuance handles book issuance creation
func (h *bookIssuanceHandler) CreateBookIssuance(c *fiber.Ctx) error {
	var issuance models.BookIssuance
	if err := c.BodyParser(&issuance); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := h.service.CreateBookIssuance(&issuance)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not issue book, db is down"})
	}

	return c.Status(201).JSON(issuance)
}

// GetAllIssuedBooks handles fetching all issued books
func (h *bookIssuanceHandler) GetAllIssuedBooks(c *fiber.Ctx) error {
	books, err := h.service.GetAllIssuedBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not find records"})
	}
	return c.Status(200).JSON(books)
}

// GetBookIssuanceByID handles fetching a single issued book by ID
func (h *bookIssuanceHandler) GetBookIssuanceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := h.service.GetBookIssuanceByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Record for the book ID not found"})
	}
	return c.Status(200).JSON(book)
}

// UpdateBookIssuance handles updating an issued book record
func (h *bookIssuanceHandler) UpdateBookIssuance(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedBook models.BookIssuance

	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	book, err := h.service.UpdateBookIssuance(id, &updatedBook)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}

// DeleteBookIssuance handles deleting an issued book record
func (h *bookIssuanceHandler) DeleteBookIssuance(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.DeleteBookIssuance(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "record not found"})
	}
	return c.JSON(fiber.Map{"message": "Record deleted."})
}
