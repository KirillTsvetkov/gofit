package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Workout struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Duration    int                `bson:"duration"`
}
