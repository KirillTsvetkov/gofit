package handler

import (
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type AuthHander struct {
	Rep *repository.Repository
}

func (h *AuthHander) singUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign up page"})
}

func (h *AuthHander) singIn(c *gin.Context) {

}
