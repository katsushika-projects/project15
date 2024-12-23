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

func GroupRoutes(r *gin.Engine, db *gorm.DB) {
	groupRepository := repositories.NewGroupRepository(db)
	userRepository := repositories.NewUserRepository(db)
	groupService := services.NewGroupService(groupRepository)
	authMiddleware := middlewares.NewAuthMiddleware(userRepository)
	groupHandler := handlers.NewGroupHandler(groupService, authMiddleware)

	groups := r.Group("/groups")
	{
		groups.POST("", groupHandler.CreateGroup)       //グループ作成
		groups.DELETE("/:id", groupHandler.DeleteGroup) //グループ削除
		groups.GET("", groupHandler.GetGroups)          //グループ一覧取得
		groups.GET("/:id", groupHandler.GetGroup)       //グループ詳細取得
	}
}

func ClassRoutes(r *gin.Engine, db *gorm.DB) {
	classes := r.Group("/classes")
	{
		classes.POST("")       //授業作成
		classes.GET("/:id")    //授業詳細取得
		classes.DELETE("/:id") //授業削除
		classes.GET("")        //授業一覧取得
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

		admin.GET("/groups", func(c *gin.Context) {
			var groups []models.Group
			if err := db.Find(&groups).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
				return
			}
			c.JSON(http.StatusOK, groups)
		})
	}
}
