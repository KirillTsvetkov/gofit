package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository, authMiddleware gin.HandlerFunc) {
	h := handler.UserHander{
		Rep: rep,
	}

	users := router.Group("/users", authMiddleware)
	{
		users.GET("/me", h.Me)
		users.GET("/:id", h.GetUserById)
		users.GET("/:id/records", h.GetRecordsUserById)
	}
}
