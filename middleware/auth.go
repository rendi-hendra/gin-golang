package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	db *gorm.DB
}

func NewAuthMiddleware(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{db: db}
}

func (am *AuthMiddleware) Auth(c *gin.Context) {
	// var user repository.User
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
		c.Abort()
		return
	}

	// am.db.Find(&user, authHeader)
	expectedToken := os.Getenv("AUTH_TOKEN")
	if expectedToken == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Server authorization missing configuration"})
		c.Abort()
		return
	}

	bearerToken := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
	bearerToken = strings.TrimSpace(strings.TrimPrefix(bearerToken, ":"))

	if bearerToken != expectedToken {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
