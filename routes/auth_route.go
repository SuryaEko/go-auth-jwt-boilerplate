package routes

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes
func AuthRoutes(router *gin.Engine, cs *controllers.ControllerService) {
	router.POST("/register", cs.Register)
	router.POST("/login", cs.Login)
}
