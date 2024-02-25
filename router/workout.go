package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/gin-gonic/gin"
)

type WorkoutRouter struct {
}

func (r *WorkoutRouter) RegisterRoutes(router *gin.Engine) {
	h := new(handler.WorkoutHandler)
	workouts := router.Group("/workouts")
	{
		workouts.GET("/", h.GetWorkouts)
		workouts.GET("/:id", h.GetWorkoutById)
	}
}
