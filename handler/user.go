package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHander struct {
}

func (h *UserHander) Me(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi"})
}

func (h *UserHander) GetUserById(c *gin.Context) {

}

func (h *UserHander) GetRecordsUserById(c *gin.Context) {

}
