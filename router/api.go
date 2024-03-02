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
	userRouter.RegisterRoutes(router, rep)

	workoutRouter := new(WorkoutRouter)
	workoutRouter.RegisterRoutes(router, rep)

	goalRouter := new(GoalRouter)
	goalRouter.RegisterRoutes(router, rep)

	achievemetRouter := new(AchievementRouter)
	achievemetRouter.RegisterRoutes(router, rep)

	return router
}
