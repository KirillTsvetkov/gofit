package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/gin-gonic/gin"
)

type GoalRouter struct {
}

func (r *GoalRouter) RegisterRoutes(router *gin.Engine) {
	h := new(handler.GoalHander)
	goals := router.Group("/goals")
	{
		goals.POST("/", h.GetGoals)
		goals.GET("/:id", h.GetGoalById)
	}
}
