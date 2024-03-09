package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserActivity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id"`
	WorkoutID primitive.ObjectID `bson:"workout_id"`
	date      time.Time          `bson:"date"`
}
