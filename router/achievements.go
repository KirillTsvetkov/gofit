package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/gin-gonic/gin"
)

type AchievementRouter struct {
}

func (r *AchievementRouter) RegisterRoutes(router *gin.Engine) {
	h := new(handler.AchievementHander)
	achievements := router.Group("/achievements")
	{
		achievements.POST("/", h.GetAchievements)
		achievements.GET("/:id", h.GetAchievementById)
	}
}
