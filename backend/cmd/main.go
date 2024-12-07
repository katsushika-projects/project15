package main

import (
	"github.com/gin-gonic/gin"
	"my-gin-app/internal/middlewares"
	"my-gin-app/internal/routes"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.SessionMiddleware("secret"))
	routes.AuthRoutes(r)
	r.Run(":8080")
}
