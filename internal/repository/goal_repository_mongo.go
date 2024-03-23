package repository

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
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

func (rep *GoalRepositoryMongo) CreateGoal(ctx context.Context, goal domain.Goal) (*domain.Goal, error) {
	res, err := rep.db.InsertOne(ctx, goal)
	if err != nil {
		return &goal, err
	}
	res.InsertedID.(primitive.ObjectID).Hex()
	return &goal, nil
}

func (rep *GoalRepositoryMongo) GetGoalById(ctx context.Context, id string) (*domain.Goal, error) {
	var goal domain.Goal
	return &goal, nil
}

func (rep *GoalRepositoryMongo) UpdateGoal(ctx context.Context, goal domain.Goal) (*domain.Goal, error) {
	filter := bson.M{"_id": goal.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if goal.UserID != primitive.NilObjectID {
		update["$set"].(bson.M)["user_id"] = goal.UserID
	}

	update["$set"].(bson.M)["status"] = goal.Status

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedGoal domain.Goal
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

func (rep *GoalRepositoryMongo) ListGoalsByUserId(ctx context.Context, user *domain.User) ([]domain.Goal, error) {
	filter := bson.M{"user_id": user.ID}
	cursor, err := rep.db.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var goals []domain.Goal
	if err = cursor.All(ctx, &goals); err != nil {
		log.Fatal(err)
	}
	return goals, nil
}
