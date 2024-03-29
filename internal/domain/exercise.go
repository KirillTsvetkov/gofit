package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ExerciseType primitive.ObjectID `json:"exerciseTypeId" bson:"exerciseTypeId"`
	Sets         int                `bson:"sets"`
	Reps         int                `bson:"reps"`
	Weight       int                `bson:"weight"`
}
