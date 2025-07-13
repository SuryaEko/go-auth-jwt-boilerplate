package routes

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes
func AuthRoutes(router *gin.Engine, controllerService *controllers.ControllerService) {
	router.POST("/register", controllerService.Register)
	router.POST("/login", controllerService.Login)
}
