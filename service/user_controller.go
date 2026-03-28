package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

func (h *UserController) FindAll(c *gin.Context) {
	h.service.FindAll(c)
}

func (h *UserController) FindByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
		return
	}
	h.service.FindByID(c, idInt)
}

func (h *UserController) Create(c *gin.Context) {
	h.service.Create(c)
}

func (h *UserController) Update(c *gin.Context) {
	h.service.Update(c)
}

func (h *UserController) Delete(c *gin.Context) {
	h.service.Delete(c)
}
