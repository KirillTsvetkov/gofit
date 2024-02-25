package repository

import (
	"github.com/KirillTsvetkov/gofit/models"
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

func (rep *AchievementRepositoryMongo) CreateAchievement(achievement models.Achievement) (models.Achievement, error) {
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) GetAchievementById(id string) (models.Achievement, error) {
	var achievement models.Achievement
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) UpdateAchievement(achievement models.Achievement) (models.Achievement, error) {
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) DeleteAchievement(id string) error {
	return nil
}

func (rep *AchievementRepositoryMongo) ListAchievementsByUserId(userId int) ([]models.Achievement, error) {
	var achievements []models.Achievement
	return achievements, nil
}

func (rep *AchievementRepositoryMongo) ListAchievementsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Achievement, error) {
	var achievements []models.Achievement
	return achievements, nil
}
