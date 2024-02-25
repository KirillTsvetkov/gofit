package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) IniteRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.GET("/sing-up", h.singIn)
		auth.GET("/sing-in", h.singUp)
	}

	users := router.Group("/users")
	{
		users.GET("/me", h.me)
		users.GET("/:id", h.getUserById)
		users.GET("/:id/records", h.getRecordsUserById)
	}

	records := router.Group("/achievements")
	{
		records.POST("/", h.getAchievements)
		records.GET("/:id", h.getAchievementById)
	}

	workouts := router.Group("/workouts")
	{
		workouts.POST("/", h.getWorkouts)
		workouts.GET("/:id", h.getWorkoutById)
	}

	return router
}
