package router

import (
	"github.com/KirillTsvetkov/gofit/internal/handler"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type WorkoutRouter struct {
}

func (r *WorkoutRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository, authMiddleware gin.HandlerFunc) {
	h := handler.WorkoutHandler{
		Rep: rep,
	}

	workouts := router.Group("/workouts", authMiddleware)
	{
		workouts.GET("/", h.GetWorkouts)
		workouts.GET("/:id", h.GetWorkoutById)
	}
}
