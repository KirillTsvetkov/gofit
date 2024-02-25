package router

import (
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) IniteRoutes(rep *repository.Repository) *gin.Engine {
	router := gin.New()

	userRouter := new(UserRouter)
	userRouter.RegisterRoutes(router)

	workoutRouter := new(WorkoutRouter)
	workoutRouter.RegisterRoutes(router)

	goalRouter := new(GoalRouter)
	goalRouter.RegisterRoutes(router)

	achievemetRouter := new(AchievementRouter)
	achievemetRouter.RegisterRoutes(router)

	return router
}
