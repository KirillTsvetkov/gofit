package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type WorkoutRouter struct {
}

func (r *WorkoutRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository) {
	h := handler.WorkoutHandler{
		Rep: rep,
	}

	workouts := router.Group("/workouts")
	{
		workouts.GET("/", h.GetWorkouts)
		workouts.GET("/:id", h.GetWorkoutById)
	}
}
