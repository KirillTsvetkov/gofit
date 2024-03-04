package router

import (
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) IniteRoutes(rep *repository.Repository, authMiddleware gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	userRouter := new(UserRouter)
	userRouter.RegisterRoutes(router, rep, authMiddleware)

	workoutRouter := new(WorkoutRouter)
	workoutRouter.RegisterRoutes(router, rep, authMiddleware)

	goalRouter := new(GoalRouter)
	goalRouter.RegisterRoutes(router, rep, authMiddleware)

	achievemetRouter := new(AchievementRouter)
	achievemetRouter.RegisterRoutes(router, rep, authMiddleware)

	return router
}
