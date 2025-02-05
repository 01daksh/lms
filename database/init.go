package database

import (
	"log"

	"lms/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	dsn := "host=localhost user=postgres password=abc.dak01 dbname=library_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Connected to the database!")

	DB.AutoMigrate(&models.BookIssuance{})
	log.Println("Database migrated!")
}
