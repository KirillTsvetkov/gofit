package router

import (
	"github.com/KirillTsvetkov/gofit/handler"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (r *UserRouter) RegisterRoutes(router *gin.Engine) {
	h := new(handler.UserHander)
	users := router.Group("/users")
	{
		users.GET("/me", h.Me)
		users.GET("/:id", h.GetUserById)
		users.GET("/:id/records", h.GetRecordsUserById)
	}
}
