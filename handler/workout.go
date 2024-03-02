package handler

import (
	"net/http"

	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	Rep *repository.Repository
}

func NewHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hi"})
}

func (h *WorkoutHandler) GetWorkoutById(c *gin.Context) {

}
