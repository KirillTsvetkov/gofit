package repository

import (
	"context"

	"github.com/KirillTsvetkov/gofit/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workout models.Workout) (*models.Workout, error)

	GetWorkoutById(ctx context.Context, id string) (*models.Workout, error)

	UpdateWorkout(workout models.Workout) (*models.Workout, error)

	DeleteWorkout(id string) error

	ListWorkoutsByUserId(userId int) ([]models.Workout, error)
}

type AchievementRepository interface {
	CreateAchievement(achievement models.Achievement) (models.Achievement, error)

	GetAchievementById(id string) (models.Achievement, error)

	UpdateAchievement(achievement models.Achievement) (models.Achievement, error)

	DeleteAchievement(id string) error

	ListAchievementsByUserId(userId int) ([]models.Achievement, error)

	ListAchievementsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Achievement, error)
}

type GoalRepository interface {
	CreateGoal(goal models.Goal) (models.Goal, error)

	GetGoalById(id string) (models.Goal, error)

	UpdateGoal(goal models.Goal) (models.Goal, error)

	DeleteGoal(id string) error

	ListGoalsByUserId(userId int) ([]models.Goal, error)

	ListGoalsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Goal, error)
}

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)

	GetUserByID(id string) (models.User, error)

	UpdateUser(user models.User) (models.User, error)

	DeleteUser(id string) error

	ListUsers() ([]models.User, error)
}

type Repository struct {
	UserRepository
	WorkoutRepository
	AchievementRepository
	GoalRepository
}

func NewRepository(dbClient *mongo.Database) *Repository {
	return &Repository{
		UserRepository:        NewUserRepositoryMongo(dbClient, viper.GetString("mongo.user_collection")),
		WorkoutRepository:     NewWorkoutRepositoryMongo(dbClient, viper.GetString("mongo.workout_collection")),
		AchievementRepository: NewAchievementRepositoryMongo(dbClient, viper.GetString("mongo.achievement_collection")),
		GoalRepository:        NewGoalRepositoryMongo(dbClient, viper.GetString("mongo.goal_collection")),
	}
}
