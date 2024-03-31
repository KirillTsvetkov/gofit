package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoalHander struct {
	service *services.GoalService
}

type GoalByRequest struct {
	Exercises []domain.Exercise `json:"exercises"`
	Date      time.Time         `json:"date"`
}

type GoalsResponse struct {
	Data []domain.Goal `json:"data"`
	Meta domain.Meta   `json:"meta"`
}

func NewGoalHander(rep *repository.Repository) *GoalHander {
	service := services.NewGoalService(rep)
	return &GoalHander{service: service}
}

func (h *GoalHander) GetGoals(c *gin.Context) {
	var request domain.GetGoalListQuery
	exerciseTypesStr, _ := c.GetQueryArray("exerciseTypes[]")

	for _, idStr := range exerciseTypesStr {
		id, err := primitive.ObjectIDFromHex(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ObjectID format"})
			return
		}
		request.ExerciseTypes = append(request.ExerciseTypes, id)
	}

	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "request": request})
		return
	}
	user := c.Value("user").(*domain.User)

	goals, total := h.service.GetUserGoals(c, user, request)
	meta := domain.Meta{Page: *request.PaginationQuery.GetPage(), Limit: *request.PaginationQuery.GetLimit(), Total: total}
	c.IndentedJSON(http.StatusOK, GoalsResponse{Data: goals, Meta: meta})
}

func (h *GoalHander) GetGoal(c *gin.Context) {
	user := c.Value("user").(*domain.User)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "Empty id param")
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid id")
		return
	}
	log.Print("before get goal")
	goal, err := h.service.GetGoal(c, user, objectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, goal)
}

func (h *GoalHander) UpdateGoal(c *gin.Context) {
	user := c.Value("user").(*domain.User)
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, "Empty id param")
		return
	}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid id")
		return
	}

	var inp domain.UpdateGoalQuery
	if err := c.BindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	goal, err := h.service.UpdateGoal(c, user, objectId, inp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, goal)
}
