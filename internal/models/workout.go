package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workout struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Exercises []Exercise         `bson:"exercises"`
	Date      time.Time          `bson:"date"`
}
