package services

import "my-gin-app/internal/models"

func Authenticate(username, password string) bool {
	user := models.GetUser(username, password)
	return user != nil
}
