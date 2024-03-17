package handler

import (
	"log"
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	service *services.WorkoutService
}

type createWorkoutRequest struct {
	Exercises []models.Exercise `json:"exercises"`
}

type WorkoutResponse struct {
	Data []models.Workout `json:"data"`
}

func NewWorkoutHandler(rep *repository.Repository) *WorkoutHandler {
	service := services.NewWorkoutService(rep)
	return &WorkoutHandler{service: service}
}

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	user := c.Value("user").(*models.User)
	workouts := h.service.GetUserWorkout(c, user)
	c.IndentedJSON(http.StatusOK, WorkoutResponse{Data: workouts})
}

func (h *WorkoutHandler) CreateWorkout(c *gin.Context) {
	user := c.Value("user").(*models.User)

	request := new(createWorkoutRequest)
	if err := c.BindJSON(request); err != nil {
		log.Print(err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	exercise := request.Exercises

	workouts := h.service.CreateWorkout(c, exercise, user)
	c.IndentedJSON(http.StatusOK, workouts)
}
