package routes

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/gin-gonic/gin"
)

// AuthRoutes sets up the authentication routes
func RegisterAllRoutes(router *gin.Engine, controllerService *controllers.ControllerService) {
	AuthRoutes(router, controllerService)
	ProfileRoutes(router, controllerService)
}
