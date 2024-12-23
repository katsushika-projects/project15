package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/middlewares"
	"github.com/moto340/project15/backend/internal/services"
)

type ClassHandler struct {
	classService   *services.ClassService
	authMiddleware *middlewares.AuthMiddleware
}

func NewClassHandler(classService *services.ClassService, authMiddleware *middlewares.AuthMiddleware) *ClassHandler {
	return &ClassHandler{classService: classService, authMiddleware: authMiddleware}
}

type ClassInput struct {
	ClassName string `json:"classname"`
	GroupID   string `json:"group_id"`
}

func (h *ClassHandler) CreateClass(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input ClassInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.classService.CreateClass(input.ClassName, input.GroupID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class Create successfully"})

}

func (h *ClassHandler) DeleteClass(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := h.classService.DeleteClass(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class delete successfully"})
}

func (h *ClassHandler) GetClass(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	class, err := h.classService.GetClass(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"class": class,
	})
}
