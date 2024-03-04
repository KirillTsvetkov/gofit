package handler

import (
	"net/http"
	"strings"

	"github.com/KirillTsvetkov/gofit/pkg/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtManager *auth.Manager
}

func NewAuthMiddleware(jwtManager *auth.Manager) gin.HandlerFunc {
	return (&AuthMiddleware{
		jwtManager: jwtManager,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authorization header is missing",
		})
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bearer token is missing",
		})
		return
	}

	if headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Bearer token is missing",
		})
		return
	}

	user, err := m.jwtManager.ValidateJWT(headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		c.AbortWithStatus(status)
		return
	}

	c.Set("user", user)
}
