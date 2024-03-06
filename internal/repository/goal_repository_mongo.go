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

type GoalRepositoryMongo struct {
	db *mongo.Collection
}

func NewGoalRepositoryMongo(dbClient *mongo.Database, collectionName string) *GoalRepositoryMongo {
	return &GoalRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *GoalRepositoryMongo) CreateGoal(ctx context.Context, goal models.Goal) (*models.Goal, error) {
	return &goal, nil
}

func (rep *GoalRepositoryMongo) GetGoalById(ctx context.Context, id string) (*models.Goal, error) {
	var goal models.Goal
	return &goal, nil
}

func (rep *GoalRepositoryMongo) UpdateGoal(ctx context.Context, goal models.Goal) (*models.Goal, error) {
	filter := bson.M{"_id": goal.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if goal.UserID != primitive.NilObjectID {
		update["$set"].(bson.M)["user_id"] = goal.UserID
	}
	if goal.WorkoutID != primitive.NilObjectID {
		update["$set"].(bson.M)["workout_id"] = goal.WorkoutID
	}
	if goal.Description != "" {
		update["$set"].(bson.M)["description"] = goal.Description
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedGoal models.Goal
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedGoal)
	if err != nil {
		return nil, err
	}

	return &updatedGoal, nil
}

func (rep *GoalRepositoryMongo) DeleteGoal(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	_, err = rep.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (rep *GoalRepositoryMongo) ListGoalsByUserId(ctx context.Context, userId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}

func (rep *GoalRepositoryMongo) ListGoalsByUserIdAndWorkoutId(ctx context.Context, userId, workoutId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}
