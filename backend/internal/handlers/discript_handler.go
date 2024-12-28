package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/middlewares"
	"github.com/moto340/project15/backend/internal/services"
)

type DiscriptHandler struct {
	discriptService *services.DiscriptService
	authMiddleware  *middlewares.AuthMiddleware
}

func NewDiscriptHandler(discriptService *services.DiscriptService, authMiddleware *middlewares.AuthMiddleware) *DiscriptHandler {
	return &DiscriptHandler{discriptService: discriptService, authMiddleware: authMiddleware}
}

type DiscriptInput struct {
	Discript string `json:"discript"`
	ClassID  string `json:"class_id"`
}

func (h *DiscriptHandler) CreateDiscript(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input DiscriptInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.discriptService.CreateDiscript(input.Discript, input.ClassID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Discript Create successfully"})
}

type DiscriptsInput struct {
	ClassID string `json:"class_id"`
}

func (h *DiscriptHandler) GetDiscripts(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input DiscriptsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	discripts, err := h.discriptService.GetDiscripts(input.ClassID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": discripts,
	})
}
