package repository

import (
	"context"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
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
	return &goal, nil
}

func (rep *GoalRepositoryMongo) DeleteGoal(ctx context.Context, id string) error {
	return nil
}

func (rep *GoalRepositoryMongo) ListGoalsByUserId(ctx context.Context, userId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}

func (rep *GoalRepositoryMongo) ListGoalsByUserIdAndWorkoutId(ctx context.Context, userId, workoutId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}
