package routes

import (
	"github.com/SuryaEko/go-auth-jwt-boilerplate/controllers"
	"github.com/SuryaEko/go-auth-jwt-boilerplate/middleware"
	"github.com/gin-gonic/gin"
)

// UserRoutes sets up the user routes
func UserRoutes(router *gin.Engine, cs *controllers.ControllerService) {
	userRouter := router.Group("/users")

	userRouter.POST("/", cs.CreateUser)
	userRouter.GET("/:id", middleware.AuthMiddleware(), cs.GetUserByID)
	userRouter.PUT("/:id", middleware.AuthMiddleware(), cs.UpdateUser)
	userRouter.PUT("/:id/password", middleware.AuthMiddleware(), cs.UpdatePasswordUser)
}
