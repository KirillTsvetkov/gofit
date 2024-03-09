package handler

import (
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	service *services.WorkoutService
}

func NewWorkoutHandler(rep *repository.Repository) *WorkoutHandler {
	service := services.NewWorkoutService(rep)
	return &WorkoutHandler{service: service}
}

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	userId := c.Value("user").(string)
	workouts := h.service.GetUserWorkout(c, userId)
	c.IndentedJSON(http.StatusOK, workouts)
}

func (h *WorkoutHandler) CreateWorkout(c *gin.Context) {
	userId := c.Value("user").(string)
	workouts := h.service.CreateWorkout(c, userId)
	c.IndentedJSON(http.StatusOK, workouts)
}
