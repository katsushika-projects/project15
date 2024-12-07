package routes

import (
	"github.com/gin-gonic/gin"
	"my-gin-app/internal/controllers"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/logout", controllers.logout)
	}
}
