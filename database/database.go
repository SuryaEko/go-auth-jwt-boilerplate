package database

import (
	"fmt"
	"log"
	"os"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect initializes a connection to the database
func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Database ping failed: %v", err)
		return nil, err
	}

	// AutoMigrate will create the User table if it doesn't exist
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
