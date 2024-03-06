package repository

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	filter := bson.M{"_id": workout.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if workout.Title != "" {
		update["$set"].(bson.M)["title"] = workout.Title
	}

	if workout.Description != "" {
		update["$set"].(bson.M)["description"] = workout.Description
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedWorkout models.Workout
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedWorkout)
	if err != nil {
		return nil, err
	}

	return &updatedWorkout, nil
}

func (rep *WorkoutRepositoryMongo) DeleteWorkout(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	_, err = rep.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (rep *WorkoutRepositoryMongo) ListWorkoutsByUserId(ctx context.Context, userId int) ([]models.Workout, error) {
	var workouts []models.Workout
	return workouts, nil
}
