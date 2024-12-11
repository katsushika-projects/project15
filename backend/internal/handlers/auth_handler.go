package handlers

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/moto340/project15/backend/internal/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
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
