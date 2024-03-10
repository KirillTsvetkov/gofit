package router

import (
	"github.com/KirillTsvetkov/gofit/internal/handler"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (r *AuthRouter) RegisterRoutes(router *gin.Engine, rep *repository.Repository, service *services.Service) {
	h := handler.NewAuthHander(service)

	goals := router.Group("/auth")
	{
		goals.POST("/token", h.GetToken)

	}
}
