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
		auth.POST("/refresh", authHandler.RefreshToken) //リフレッシュトークンの検証
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
		groups.POST("/get", groupHandler.GetGroups)     //グループ一覧取得
		groups.GET("/:id", groupHandler.GetGroup)       //グループ詳細取得
	}
}

func ClassRoutes(r *gin.Engine, db *gorm.DB) {
	classRepository := repositories.NewClassRepository(db)
	userRepository := repositories.NewUserRepository(db)
	classService := services.NewClassService(classRepository)
	authMiddleware := middlewares.NewAuthMiddleware(userRepository)
	classHandler := handlers.NewClassHandler(classService, authMiddleware)

	classes := r.Group("/classes")
	{
		classes.POST("", classHandler.CreateClass)       //授業作成
		classes.GET("/:id", classHandler.GetClass)       //授業詳細取得
		classes.DELETE("/:id", classHandler.DeleteClass) //授業削除
		classes.POST("/get", classHandler.GetClasses)    //授業一覧取得
	}
}

func Discription(r *gin.Engine, db *gorm.DB) {
	discriptRepository := repositories.NewDiscriptRepository(db)
	userRepository := repositories.NewUserRepository(db)
	discriptService := services.NewDiscriptService(discriptRepository)
	authMiddleware := middlewares.NewAuthMiddleware(userRepository)
	discriptHandler := handlers.NewDiscriptHandler(discriptService, authMiddleware)

	discription := r.Group("/discription")
	{
		discription.POST("", discriptHandler.CreateDiscript) //質問、回答投稿
		discription.GET("", discriptHandler.GetDiscripts)    //スレ一覧取得
		discription.POST("/:id")                             //スレ削除
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
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch groups"})
				return
			}
			c.JSON(http.StatusOK, groups)
		})

		admin.GET("/classes", func(c *gin.Context) {
			var classes []models.Class
			if err := db.Find(&classes).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
				return
			}
			c.JSON(http.StatusOK, classes)
		})

		admin.GET("/discription", func(c *gin.Context) {
			var discription []models.Thread
			if err := db.Find(&discription).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch discription"})
				return
			}
			c.JSON(http.StatusOK, discription)
		})
	}
}
