package repository

import (
	"github.com/KirillTsvetkov/gofit/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRepositoryMongo struct {
	dbClient *mongo.Client
}

func NewWorkoutRepositoryMongo(dbClient *mongo.Client, collectionName string) *WorkoutRepositoryMongo {
	return &WorkoutRepositoryMongo{
		dbClient: dbClient,
	}
}

func (rep *WorkoutRepositoryMongo) CreateWorkout(workout models.Workout) (string, error) {
	return "temp", nil
}

func (rep *WorkoutRepositoryMongo) GetWorkoutById(id string) (models.Workout, error) {
	var workout models.Workout
	return workout, nil
}

func (rep *WorkoutRepositoryMongo) UpdateWorkout(workout models.Workout) error {
	return nil
}

func (rep *WorkoutRepositoryMongo) DeleteWorkout(id string) error {
	return nil
}

func (rep *WorkoutRepositoryMongo) ListWorkoutsByUserId(userId int) ([]models.Workout, error) {
	var workouts []models.Workout
	return workouts, nil
}
