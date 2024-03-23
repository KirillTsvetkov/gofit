package handler

import (
	"log"
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	service *services.WorkoutService
}

type CreateWorkoutRequest struct {
	Exercises []domain.Exercise `json:"exercises"`
}

type Meta struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}

type WorkoutResponse struct {
	Data []domain.Workout `json:"data"`
	Meta Meta             `json:"meta"`
}

func NewWorkoutHandler(rep *repository.Repository) *WorkoutHandler {
	service := services.NewWorkoutService(rep)
	return &WorkoutHandler{service: service}
}

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	var request domain.GetWorkoutListQuery
	log.Print("query: ", c.Request.URL.Query())

	if err := c.Bind(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("user").(*domain.User)

	workouts, total := h.service.GetUserWorkout(c, user, request)
	meta := Meta{Page: request.Page, Limit: request.Limit, Total: total}
	c.IndentedJSON(http.StatusOK, WorkoutResponse{Data: workouts, Meta: meta})
}

func (h *WorkoutHandler) CreateWorkout(c *gin.Context) {
	user := c.Value("user").(*domain.User)

	request := new(CreateWorkoutRequest)
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	exercise := request.Exercises
	if len(exercise) < 1 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	workouts := h.service.CreateWorkout(c, exercise, user)
	c.IndentedJSON(http.StatusOK, workouts)
}
