package routes

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

// ProfileRoutes sets up the profile routes
func ProfileRoutes(router *gin.Engine, controllerService *controllers.ControllerService) {
	router.GET("/profile", middleware.AuthMiddleware(), controllerService.Profile)
}
