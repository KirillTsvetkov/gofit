package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	Exercises []Exercise         `bson:"exercises"`
	Date      time.Time          `bson:"date"`
}

type WorkoutFilterQuery struct {
	DateFrom time.Time `form:"dateFrom" json:"dateFrom" time_format:"2006-01-02"`
	DateTo   time.Time `form:"dateTo" json:"dateTo" time_format:"2006-01-02"`
}

type GetWorkoutListQuery struct {
	PaginationQuery
	WorkoutFilterQuery
}

type UpdateWorkoutQuery struct {
	Exercises []Exercise `bson:"exercises"`
	Date      time.Time  `bson:"date"`
}
