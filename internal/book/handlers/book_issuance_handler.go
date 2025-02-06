package handlers

import (
	"lms/internal/Book/service"
	"lms/models"

	"github.com/gofiber/fiber/v2"
)

func CreateBookIssuance(c *fiber.Ctx) error {
	var issuance models.BookIssuance
	if err := c.BodyParser(&issuance); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	err := service.CreateBookIssuance(&issuance)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not issue book, db is down"})
	}

	return c.Status(201).JSON(issuance)
}

func GetAllIssuedBooks(c *fiber.Ctx) error {
	books, err := service.GetAllIssuedBooks()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not find"})
	}
	return c.Status(200).JSON(books)
}

func GetBookIssuanceByID(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := service.GetBookIssuanceByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Record for the book id not found"})
	}
	return c.Status(200).JSON(book)
}

func UpdateBookIssuance(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedBook models.BookIssuance

	if err := c.BodyParser(&updatedBook); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	book, err := service.UpdateBookIssuance(id, &updatedBook)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(book)
}

// DeleteBookIssuance deletes an issued book record
func DeleteBookIssuance(c *fiber.Ctx) error {
	id := c.Params("id")

	err := service.DeleteBookIssuance(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "record not found"})
	}
	return c.JSON(fiber.Map{"message": "Record deleted."})
}
