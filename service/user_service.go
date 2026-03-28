package service

import "github.com/gin-gonic/gin"

type UserService interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindByID(c *gin.Context, id int)
	FindAll(c *gin.Context)
}
