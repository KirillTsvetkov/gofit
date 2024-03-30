package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Goal struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserID   primitive.ObjectID `bson:"user_id"`
	Exercise Exercise           `bson:"exercise"`
	Date     time.Time          `bson:"date"`
	Status   GoalStatus         `bson:"status"`
}

type GoalFilterQuery struct {
	DateFrom      time.Time            `form:"dateFrom" json:"dateFrom" time_format:"2006-01-02"`
	DateTo        time.Time            `form:"dateTo" json:"dateTo" time_format:"2006-01-02"`
	ExerciseTypes []primitive.ObjectID `binding:"-"`
	Status        GoalStatus           `form:"GoalStatus" json:"GoalStatus"`
}
type GetGoalListQuery struct {
	PaginationQuery
	GoalFilterQuery
}
