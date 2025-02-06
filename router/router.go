package router

import (
	"lms/internal/book/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/issue", handlers.CreateBookIssuance)
	app.Get("/issued-books", handlers.GetAllIssuedBooks)
	app.Get("/issued-books/:id", handlers.GetBookIssuanceByID)
	app.Put("/issued-books/:id", handlers.UpdateBookIssuance)
	app.Delete("/issued-books/:id", handlers.DeleteBookIssuance)
}
