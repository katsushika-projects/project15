package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/middlewares"
	"github.com/moto340/project15/backend/internal/services"
)

type GroupHandler struct {
	groupService   *services.GroupService
	authMiddleware *middlewares.AuthMiddleware
}

func NewGroupHandler(groupService *services.GroupService, authMiddleware *middlewares.AuthMiddleware) *GroupHandler {
	return &GroupHandler{groupService: groupService, authMiddleware: authMiddleware}
}

type GroupInput struct {
	University string `json:"university"`
	Fculty     string `json:"fculty"`
	Department string `json:"department"`
	Grade      string `json:"grade"`
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input GroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.groupService.CreateGroup(input.University, input.Fculty, input.Department, input.Grade); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Geoup Create successfully"})

}

func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := h.groupService.DeleteGroup(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Geoup Delete successfully"})
}

func (h *GroupHandler) GetGroups(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if err := h.authMiddleware.AuthAccessToken(authHeader); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input GroupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	groups, err := h.groupService.GetGroups(input.University, input.Fculty, input.Department, input.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}
