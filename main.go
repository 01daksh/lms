package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
struct contains (var_name, datatype, struct_tags)
struct tags here helps in:

	identifying json from api request
*/
type BookIssuance struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	BookTitle  string `json:"book_title"`
	IssuedTo   string `json:"issued_to"`
	IssuedDate string `json:"issued_date"`
	ReturnDate string `json:"return_date"`
}

var db *gorm.DB

// gorm is used for connecting to postgres in go lang

func initDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=abc.dak01 dbname=library_db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connected to the database!")

	db.AutoMigrate(&BookIssuance{}) // it is used for creating the database schema automatically
	log.Println("Database migrated!")
}

func main() {
	app := fiber.New()

	initDatabase()

	// fmt.Println(app)

	app.Post("/issue", func(c *fiber.Ctx) error {
		var issuance BookIssuance
		err := c.BodyParser(&issuance)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "invalid body"})
		}
		db.Create(&issuance)
		return c.Status(201).JSON(issuance)
	})

	app.Get("/issued-books", func(c *fiber.Ctx) error {
		books := []BookIssuance{}
		db.Find(&books)
		return c.Status(200).JSON(books)

	})

	app.Get("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		// book := BookIssuance (cant' be used becuase the shorthhand expects an initialisation on rhs)
		var book BookIssuance
		err := db.First(&book, id).Error
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Record for the book id not found"})
		}

		return c.Status(200).JSON(book)
	})

	app.Put("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book BookIssuance

		if err := db.First(&book, id).Error; err != nil {
			return c.Status(404).JSON("No Result Found")
		}
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON("invalid body.")
		}

		db.Save(&book)
		return c.JSON(book)
	})

	app.Delete("/issued-books/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var book BookIssuance
		if err := db.First(&book, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Record Not Found"})
		}

		db.Delete(&book)
		return c.JSON(fiber.Map{"message": "Record deleted."})

	})

	app.Listen(":3000")

}
