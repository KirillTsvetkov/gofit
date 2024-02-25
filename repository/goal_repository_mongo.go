package repository

import (
	"github.com/KirillTsvetkov/gofit/models"
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

func (rep *GoalRepositoryMongo) CreateGoal(goal models.Goal) (models.Goal, error) {
	return goal, nil
}

func (rep *GoalRepositoryMongo) GetGoalById(id string) (models.Goal, error) {
	var goal models.Goal
	return goal, nil
}

func (rep *GoalRepositoryMongo) UpdateGoal(goal models.Goal) (models.Goal, error) {
	return goal, nil
}

func (rep *GoalRepositoryMongo) DeleteGoal(id string) error {
	return nil
}

func (rep *GoalRepositoryMongo) ListGoalsByUserId(userId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}

func (rep *GoalRepositoryMongo) ListGoalsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Goal, error) {
	var models []models.Goal
	return models, nil
}
