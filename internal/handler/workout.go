package handler

import (
	"log"
	"net/http"
	"time"

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
	Date      time.Time         `json:"date"`
}

type WorkoutResponse struct {
	Data []domain.Workout `json:"data"`
	Meta domain.Meta      `json:"meta"`
}

func NewWorkoutHandler(rep *repository.Repository) *WorkoutHandler {
	service := services.NewWorkoutService(rep)
	return &WorkoutHandler{service: service}
}

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	var request domain.GetWorkoutListQuery

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Print(request)
	user := c.Value("user").(*domain.User)

	workouts, total := h.service.GetUserWorkout(c, user, request)
	meta := domain.Meta{Page: *request.PaginationQuery.GetPage(), Limit: *request.PaginationQuery.GetLimit(), Total: total}
	c.IndentedJSON(http.StatusOK, WorkoutResponse{Data: workouts, Meta: meta})
}

func (h *WorkoutHandler) CreateWorkout(c *gin.Context) {
	user := c.Value("user").(*domain.User)

	request := new(CreateWorkoutRequest)
	if err := c.BindJSON(request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(request.Exercises) < 1 {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	workouts := h.service.CreateWorkout(c, request.Date, request.Exercises, user)
	c.IndentedJSON(http.StatusOK, workouts)
}
