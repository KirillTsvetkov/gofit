package repository

import (
	"github.com/KirillTsvetkov/gofit/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AchievementRepositoryMongo struct {
	dbClient *mongo.Client
}

func NewAchievementRepositoryMongo(dbClient *mongo.Client) *AchievementRepositoryMongo {
	return &AchievementRepositoryMongo{
		dbClient: dbClient,
	}
}

func (rep *AchievementRepositoryMongo) CreateAchievement(achievement models.Achievement) (string, error) {
	return "temp", nil
}

func (rep *AchievementRepositoryMongo) GetAchievementById(id string) (models.Achievement, error) {
	var achievement models.Achievement
	return achievement, nil
}

func (rep *AchievementRepositoryMongo) UpdateAchievement(achievement models.Achievement) error {
	return nil
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
