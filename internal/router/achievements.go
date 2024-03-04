package router

import (
	"github.com/KirillTsvetkov/gofit/internal/handler"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type AchievementRouter struct {
}

func (r *AchievementRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository, authMiddleware gin.HandlerFunc) {
	h := handler.AchievementHander{
		Rep: rep,
	}

	achievements := router.Group("/achievements", authMiddleware)
	{
		achievements.POST("/", h.GetAll)
		achievements.GET("/:id", h.Get)
		achievements.DELETE("/:id", h.Delete)
	}
}
