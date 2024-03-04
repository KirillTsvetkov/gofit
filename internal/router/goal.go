package router

import (
	"github.com/KirillTsvetkov/gofit/internal/handler"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type GoalRouter struct {
}

func (r *GoalRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository, authMiddleware gin.HandlerFunc) {
	h := handler.GoalHander{
		Rep: rep,
	}

	goals := router.Group("/goals", authMiddleware)
	{
		goals.POST("/", h.GetGoals)
		goals.GET("/:id", h.GetGoalById)
	}
}
