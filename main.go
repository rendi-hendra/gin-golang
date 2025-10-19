package main

import (
	"gin-golang/database"
	"gin-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	router := gin.Default()
	router.GET("/users", service.GetUsers)
	router.GET("/users/:id", service.GetUserByID)
	router.POST("/users", service.CreateUsers)

	router.Run("localhost:8080")
}
