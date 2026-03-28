package main

import (
	"gin-golang/database"
	"gin-golang/middleware"
	"gin-golang/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	userService := service.NewUserServiceImpl(database.DB)
	userController := service.NewUserController(userService)

	router := gin.Default()
	user := router.Group("/api/users")
	user.Use(middleware.NewAuthMiddleware(database.DB).Auth)
	user.GET("", userController.FindAll)
	user.GET("/:id", userController.FindByID)
	user.POST("", userController.Create)
	user.PUT("/:id", userController.Update)
	user.DELETE("/:id", userController.Delete)

	router.Run("localhost:8080")
}
