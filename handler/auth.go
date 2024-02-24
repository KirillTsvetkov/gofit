package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Sign up page"})
}

func (h *Handler) singIn(c *gin.Context) {

}
