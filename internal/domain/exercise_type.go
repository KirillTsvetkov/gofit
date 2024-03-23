package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseType struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}
