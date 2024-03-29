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

type AchievementRepositoryMongo struct {
	db *mongo.Collection
}

func NewAchievementRepositoryMongo(dbClient *mongo.Database, collectionName string) *AchievementRepositoryMongo {
	return &AchievementRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *AchievementRepositoryMongo) CreateAchievement(ctx context.Context, achievement domain.Achievement) (domain.Achievement, error) {
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) GetAchievementById(ctx context.Context, id string) (*domain.Achievement, error) {
	var achievement domain.Achievement
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}
	result := rep.db.FindOne(ctx, bson.M{"_id": objectId})

	if err := result.Decode(&achievement); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &achievement, nil
}

func (rep *AchievementRepositoryMongo) UpdateAchievement(ctx context.Context, achievement domain.Achievement) (*domain.Achievement, error) {
	filter := bson.M{"_id": achievement.ID}
	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	if achievement.UserID != primitive.NilObjectID {
		update["$set"].(bson.M)["user_id"] = achievement.UserID
	}
	if achievement.WorkoutID != primitive.NilObjectID {
		update["$set"].(bson.M)["workout_id"] = achievement.WorkoutID
	}
	if achievement.Description != "" {
		update["$set"].(bson.M)["description"] = achievement.Description
	}

	findUpdateOptions := options.FindOneAndUpdateOptions{}
	findUpdateOptions.SetReturnDocument(options.After)

	var updatedAchievement domain.Achievement
	err := rep.db.FindOneAndUpdate(ctx, filter, update, &findUpdateOptions).Decode(&updatedAchievement)
	if err != nil {
		return nil, err
	}

	return &updatedAchievement, nil
}

func (rep *AchievementRepositoryMongo) DeleteAchievement(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	_, err = rep.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (rep *AchievementRepositoryMongo) ListAchievements(ctx context.Context, user *domain.User, pagination domain.PaginationQuery) ([]domain.Achievement, int64, error) {
	filter := bson.M{"user_id": user.ID}

	totalCount, err := rep.db.CountDocuments(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fOpt := pagination.GetPaginationOpts()

	cursor, err := rep.db.Find(ctx, filter, fOpt)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var achievements []domain.Achievement
	for cursor.Next(ctx) {
		var achievement domain.Achievement
		if err := cursor.Decode(&achievement); err != nil {
			log.Println(err)
		}

		achievements = append(achievements, achievement)
	}

	return achievements, totalCount, nil
}
