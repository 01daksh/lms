package router

import (
	"lms/models"
	"lms/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/issue", func(c *fiber.Ctx) error {
		var issuance models.BookIssuance
		err := c.BodyParser(&issuance)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "invalid body"})
		}
		err = service.CreateBookIssuance(&issuance)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "could not issue book, db is down"})
		}
		return c.Status(201).JSON(issuance)
	})

	app.Get("/issued-books", func(c *fiber.Ctx) error {
		books, err := service.GetAllIssuedBooks() // this return 2 values
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "could not find"})
		}
		return c.Status(200).JSON(books)
	})

	app.Get("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		book, err := service.GetBookIssuanceByID(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Record for the book id not found"})
		}
		return c.Status(200).JSON(book)
	})

	app.Put("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var updatedBook models.BookIssuance

		if err := c.BodyParser(&updatedBook); err != nil {
			return c.Status(400).JSON("invalid body.")
		}
		book, err := service.UpdateBookIssuance(id, &updatedBook)

		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(book)
	})

	app.Delete("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := service.DeleteBookIssuance(id)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "record not found"})
		}

		return c.JSON(fiber.Map{"message": "Record deleted."})
	})
}
