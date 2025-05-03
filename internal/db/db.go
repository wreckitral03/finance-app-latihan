package db

import (
	"fmt"
	"log"
	"os"

	"finance-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Get environment variables for DB connection
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create DSN using environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// Open the database connection
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run migrations on the Transaction model
	log.Println("Connecting to DB:", dsn)
	err = database.AutoMigrate(&models.Transaction{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	} else {
		log.Println("Migration successful.")
	}

	// Set the global DB variable
	DB = database
}
