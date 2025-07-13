package main

import (
	"fmt"
	"log"

	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/database"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/routes"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Test database connection
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	services := services.InitServiceContainer(db)

	controllerService := &controllers.ControllerService{
		Services: services,
	}

	router := gin.Default()
	routes.RegisterAllRoutes(router, controllerService)

	// Start the server
	router.Run(":8080")

	fmt.Println("Server is running on port 8080")
}
