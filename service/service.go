package service

import (
	"fmt"
	"gin-golang/database"
	"gin-golang/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []repository.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func CreateUsers(c *gin.Context) {
	var newUser repository.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// repository.Users = append(repository.Users, newUser)
	user := database.DB.Create(&newUser)
	if user.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": user.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func GetUserByID(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.ParseUint(param, 10, 32)
	if err != nil {
		fmt.Printf("Error converting string to uint32: %v\n", err)
		return
	}

	for _, a := range repository.Users {
		if a.ID == uint(id) {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
