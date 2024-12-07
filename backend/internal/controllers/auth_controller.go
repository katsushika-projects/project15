package controllers

import (
	"net/http"

	"my-gin-app/internal/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/moto340/project15/backend/internal/services"
)

func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !services.Authenticate(credentials.Username, credentials.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("username", credentials.Username)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func Logout(c *gin.Context) {
	session := session.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
