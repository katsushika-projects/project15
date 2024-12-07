package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"my-gin-app/internal/handlers"
	"my-gin-app/internal/repositories"
	"my-gin-app/internal/services"
)

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/signup", authHandler.Signup)
	}
}
