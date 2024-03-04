package handler

import (
	"net/http"

	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/gin-gonic/gin"
)

type AchievementHander struct {
	Rep *repository.Repository
}

func (h *AchievementHander) GetAll(c *gin.Context) {

}

func (h *AchievementHander) Get(c *gin.Context) {
	id := c.Param("id")

	achievement, err := h.Rep.AchievementRepository.GetAchievementById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, achievement)
}

func (h *AchievementHander) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.Rep.AchievementRepository.DeleteAchievement(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}
