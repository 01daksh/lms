package main

import (
	"lms/database"
	"lms/router"

	"github.com/gofiber/fiber/v2"
)

/*
struct contains (var_name, datatype, struct_tags)
struct tags here helps in:

	identifying json from api request
*/

func main() {
	app := fiber.New()

	database.InitDatabase()

	// fmt.Println(app)

	router.SetupRoutes(app)

	app.Listen(":3000")

}
