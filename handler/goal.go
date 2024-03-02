package handler

import (
	"github.com/KirillTsvetkov/gofit/repository"
	"github.com/gin-gonic/gin"
)

type GoalHander struct {
	Rep *repository.Repository
}

func (h *GoalHander) Me(c *gin.Context) {

}

func (h *GoalHander) GetGoals(c *gin.Context) {

}

func (h *GoalHander) GetGoalById(c *gin.Context) {

}
