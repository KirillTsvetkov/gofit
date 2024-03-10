package router

import (
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) IniteRoutes(rep *repository.Repository, service *services.Service, authMiddleware gin.HandlerFunc) *gin.Engine {
	router := gin.New()

	authRouter := new(AuthRouter)
	authRouter.RegisterRoutes(router, rep, service)

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
