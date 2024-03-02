package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type GoalRouter struct {
}

func (r *GoalRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository) {
	h := handler.GoalHander{
		Rep: rep,
	}

	goals := router.Group("/goals")
	{
		goals.POST("/", h.GetGoals)
		goals.GET("/:id", h.GetGoalById)
	}
}
