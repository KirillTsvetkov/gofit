package repository

import (
	"context"

	"github.com/KirillTsvetkov/gofit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRepositoryMongo struct {
	db *mongo.Collection
}

func NewWorkoutRepositoryMongo(dbClient *mongo.Database, collectionName string) *WorkoutRepositoryMongo {
	return &WorkoutRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *WorkoutRepositoryMongo) CreateWorkout(ctx context.Context, workout models.Workout) (*models.Workout, error) {
	res, err := rep.db.InsertOne(ctx, workout)
	if err != nil {
		return &workout, err
	}
	res.InsertedID.(primitive.ObjectID).Hex()
	return &workout, nil
}

func (rep *WorkoutRepositoryMongo) GetWorkoutById(ctx context.Context, id string) (*models.Workout, error) {
	workout := new(models.Workout)
	err := rep.db.FindOne(ctx, bson.M{
		"ID": id,
	}).Decode(workout)

	if err != nil {
		return nil, err
	}
	return workout, nil
}

func (rep *WorkoutRepositoryMongo) UpdateWorkout(ctx context.Context, workout models.Workout) (*models.Workout, error) {
	return &workout, nil
}

func (rep *WorkoutRepositoryMongo) DeleteWorkout(ctx context.Context, id string) error {
	return nil
}

func (rep *WorkoutRepositoryMongo) ListWorkoutsByUserId(ctx context.Context, userId int) ([]models.Workout, error) {
	var workouts []models.Workout
	return workouts, nil
}
