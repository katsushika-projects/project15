package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/moto340/project15/backend/internal/middlewares"
	"github.com/moto340/project15/backend/internal/services"
)

type AuthHandler struct {
	authService    *services.AuthService
	authMiddleware *middlewares.AuthMiddleware
}

func NewAuthHandler(authService *services.AuthService, authMiddleware *middlewares.AuthMiddleware) *AuthHandler {
	return &AuthHandler{authService: authService, authMiddleware: authMiddleware}
}

type SignupInput struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,min=6"`
}

func validateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	// ユーザー名は英数字3〜20文字のみ許可
	re := regexp.MustCompile(`^[a-zA-Z0-9]{3,20}$`)
	return re.MatchString(username)
}

func (h *AuthHandler) Signup(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// このリクエストのバリデーションに対してのみカスタムルールを登録
		v.RegisterValidation("username", validateUsername)
	}

	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Signup(input.Username, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type LoginInput struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,min=6"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// このリクエストのバリデーションに対してのみカスタムルールを登録
		v.RegisterValidation("username", validateUsername)
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.Login(input.Username, input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.authService.GenerateTokens(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.RefreshTokenDisable(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.AccessTokenDisable(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Logout successfully"})

}
