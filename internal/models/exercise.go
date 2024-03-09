package models

type Exercise struct {
	ExerciseType ExerciseType `bson:"exerciseType"`
	Sets         int          `bson:"sets"`
	Reps         int          `bson:"reps"`
	Weight       int          `bson:"weight"`
}
