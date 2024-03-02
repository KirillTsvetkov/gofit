package repository

import (
	"context"
	"log"

	"github.com/KirillTsvetkov/gofit/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AchievementRepositoryMongo struct {
	db *mongo.Collection
}

func NewAchievementRepositoryMongo(dbClient *mongo.Database, collectionName string) *AchievementRepositoryMongo {
	return &AchievementRepositoryMongo{
		db: dbClient.Collection(collectionName),
	}
}

func (rep *AchievementRepositoryMongo) CreateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error) {
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) GetAchievementById(ctx context.Context, id string) (*models.Achievement, error) {
	var achievement models.Achievement
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

func (rep *AchievementRepositoryMongo) UpdateAchievement(ctx context.Context, achievement models.Achievement) (*models.Achievement, error) {
	return &achievement, nil
}

func (rep *AchievementRepositoryMongo) DeleteAchievement(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Invalid id")
	}

	_, err = rep.db.DeleteOne(ctx, bson.M{"_id": objectId})
	return err
}

func (rep *AchievementRepositoryMongo) ListAchievementsByUserId(ctx context.Context, userId int) ([]models.Achievement, error) {
	var achievements []models.Achievement
	return achievements, nil
}

func (rep *AchievementRepositoryMongo) ListAchievementsByUserIdAndWorkoutId(ctx context.Context, userId, workoutId int) ([]models.Achievement, error) {
	var achievements []models.Achievement
	return achievements, nil
}
