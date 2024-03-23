package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Achievement struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	WorkoutID   primitive.ObjectID `bson:"workout_id"`
	Description string             `bson:"description"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}
