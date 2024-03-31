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

func (rep *GoalRepositoryMongo) GetGoalById(ctx context.Context, user *domain.User, id primitive.ObjectID) (domain.Goal, error) {
	var goal domain.Goal
	filter := bson.M{"_id": id, "user_id": user.ID}
	if err := rep.db.FindOne(ctx, filter).Decode(&goal); err != nil {
		log.Println(err)
		return goal, err
	}
	return goal, nil
}

func (rep *GoalRepositoryMongo) UpdateGoal(ctx context.Context, goal domain.Goal, query domain.UpdateGoalQuery) (domain.Goal, error) {
	filter := bson.M{"_id": goal.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if !query.Date.IsZero() {
		update["$set"].(bson.M)["date"] = query.Date
	}

	if query.Exercise != (domain.Exercise{}) {
		update["$set"].(bson.M)["exercise"] = query.Exercise
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedGoal domain.Goal
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedGoal)
	if err != nil {
		return updatedGoal, err
	}

	return updatedGoal, nil
}

func (rep *GoalRepositoryMongo) DeleteGoal(ctx context.Context, id primitive.ObjectID) error {
	_, err := rep.db.DeleteOne(ctx, bson.M{"_id": id})
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

func (rep *GoalRepositoryMongo) ListGoals(ctx context.Context, user *domain.User, query domain.GetGoalListQuery) ([]domain.Goal, int64, error) {
	filter := bson.M{"user_id": user.ID}

	if len(query.ExerciseTypes) != 0 {
		filter["$and"] = []bson.M{}

		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			"exercise.exerciseTypeId": bson.M{"$in": query.ExerciseTypes},
		})
		log.Print(filter)
	}

	if err := filterDateQueries(query.GoalFilterQuery.DateFrom, query.GoalFilterQuery.DateTo, "date", filter); err != nil {
		return nil, 0, err
	}

	totalCount, err := rep.db.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fOpt := query.PaginationQuery.GetPaginationOpts()

	cursor, err := rep.db.Find(ctx, filter, fOpt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var goals []domain.Goal
	for cursor.Next(ctx) {
		var goal domain.Goal
		if err := cursor.Decode(&goal); err != nil {
			log.Println(err)
		}

		goals = append(goals, goal)
	}

	return goals, totalCount, nil
}
