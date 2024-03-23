package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Exercise struct {
	ExerciseType primitive.ObjectID `bson:"exerciseTypeId"`
	Sets         int                `bson:"sets"`
	Reps         int                `bson:"reps"`
	Weight       int                `bson:"weight"`
}
