package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi"})
}

func (h *Handler) getUserById(c *gin.Context) {

}

func (h *Handler) getRecordsUserById(c *gin.Context) {

}
