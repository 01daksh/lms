package router

import (
	"lms/database"
	"lms/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/issue", func(c *fiber.Ctx) error {
		var issuance models.BookIssuance
		err := c.BodyParser(&issuance)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "invalid body"})
		}
		database.DB.Create(&issuance)
		return c.Status(201).JSON(issuance)
	})

	app.Get("/issued-books", func(c *fiber.Ctx) error {
		books := []models.BookIssuance{}
		database.DB.Find(&books)
		return c.Status(200).JSON(books)
	})

	app.Get("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.BookIssuance
		err := database.DB.First(&book, id).Error
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Record for the book id not found"})
		}
		return c.Status(200).JSON(book)
	})

	app.Put("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.BookIssuance

		if err := database.DB.First(&book, id).Error; err != nil {
			return c.Status(404).JSON("No Result Found")
		}
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON("invalid body.")
		}

		database.DB.Save(&book)
		return c.JSON(book)
	})

	app.Delete("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book models.BookIssuance
		if err := database.DB.First(&book, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Record Not Found"})
		}

		database.DB.Delete(&book)
		return c.JSON(fiber.Map{"message": "Record deleted."})
	})
}
