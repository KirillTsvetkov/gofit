package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Goal struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	WorkoutID   primitive.ObjectID `bson:"workout_id"`
	Description string             `bson:"description"`
	GoalDate    time.Time          `bson:"goal_date"`
}
