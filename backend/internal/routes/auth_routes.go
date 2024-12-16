package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/handlers"
	"github.com/moto340/project15/backend/internal/middlewares"
	"github.com/moto340/project15/backend/internal/models"
	"github.com/moto340/project15/backend/internal/repositories"
	"github.com/moto340/project15/backend/internal/services"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authMiddleware := middlewares.NewAuthMiddleware(userRepository)
	authHandler := handlers.NewAuthHandler(authService, authMiddleware)

	r.POST("/register", authHandler.Signup)
}

func AuthRoutes(r *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authMiddleware := middlewares.NewAuthMiddleware(userRepository)
	authHandler := handlers.NewAuthHandler(authService, authMiddleware)

	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}
}

func AdminRoutes(r *gin.Engine, db *gorm.DB) {
	admin := r.Group("/admin")
	{
		admin.GET("/users", func(c *gin.Context) {
			var users []models.User
			if err := db.Find(&users).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
				return
			}
			c.JSON(http.StatusOK, users)
		})
	}
}
