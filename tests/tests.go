package tests

import (
	"log"
	"os"
	"testing"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/database"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/models"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/routes"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter(cs *controllers.ControllerService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	routes.RegisterAllRoutes(r, cs)
	return r
}

func setupIntegrationTest(t *testing.T) (*gin.Engine, error) {
	// Load environment variables from .env.test file
	godotenv.Load("../.env.test")

	dbHost := os.Getenv("DB_HOST")
	log.Printf("Using database host: %s", dbHost)

	db, err := database.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.Migrator().DropTable(&models.User{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	// Mock service container & controller (replace with real/mock DB as needed)
	// For real test, use a test DB or mock DB
	serviceContainer := services.InitServiceContainer(db)

	cs := &controllers.ControllerService{Services: serviceContainer}
	r := setupRouter(cs)

	return r, nil
}
