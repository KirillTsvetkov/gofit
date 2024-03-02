package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type AchievementRouter struct {
}

func (r *AchievementRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository) {
	h := handler.AchievementHander{
		Rep: rep,
	}

	achievements := router.Group("/achievements")
	{
		achievements.POST("/", h.GetAll)
		achievements.GET("/:id", h.Get)
		achievements.DELETE("/:id", h.Delete)
	}
}
