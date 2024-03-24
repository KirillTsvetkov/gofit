package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExerciseType struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}
