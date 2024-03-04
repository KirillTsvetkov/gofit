package repository

import (
	"context"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workout models.Workout) (*models.Workout, error)

	GetWorkoutById(ctx context.Context, id string) (*models.Workout, error)

	UpdateWorkout(ctx context.Context, workout models.Workout) (*models.Workout, error)

	DeleteWorkout(ctx context.Context, id string) error

	ListWorkoutsByUserId(ctx context.Context, userId int) ([]models.Workout, error)
}

type AchievementRepository interface {
	CreateAchievement(ctx context.Context, achievement models.Achievement) (models.Achievement, error)

	GetAchievementById(ctx context.Context, id string) (*models.Achievement, error)

	UpdateAchievement(ctx context.Context, achievement models.Achievement) (*models.Achievement, error)

	DeleteAchievement(ctx context.Context, id string) error

	ListAchievementsByUserId(ctx context.Context, userId int) ([]models.Achievement, error)

	ListAchievementsByUserIdAndWorkoutId(ctx context.Context, userId, workoutId int) ([]models.Achievement, error)
}

type GoalRepository interface {
	CreateGoal(ctx context.Context, goal models.Goal) (*models.Goal, error)

	GetGoalById(ctx context.Context, id string) (*models.Goal, error)

	UpdateGoal(ctx context.Context, goal models.Goal) (*models.Goal, error)

	DeleteGoal(ctx context.Context, id string) error

	ListGoalsByUserId(ctx context.Context, userId int) ([]models.Goal, error)

	ListGoalsByUserIdAndWorkoutId(ctx context.Context, userId, workoutId int) ([]models.Goal, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) (*models.User, error)

	GetUserByID(ctx context.Context, id string) (*models.User, error)

	UpdateUser(ctx context.Context, user models.User) (*models.User, error)

	DeleteUser(ctx context.Context, id string) error

	ListUsers(ctx context.Context) ([]models.User, error)
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
