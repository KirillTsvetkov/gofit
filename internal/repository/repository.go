package repository

import (
	"context"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkoutRepository interface {
	CreateWorkout(ctx context.Context, workout *domain.Workout, user *domain.User) (*domain.Workout, error)

	GetWorkoutById(ctx context.Context, id string) (*domain.Workout, error)

	GetWorkoutByUserId(ctx context.Context, user *domain.User) ([]domain.Workout, error)

	UpdateWorkout(ctx context.Context, workout domain.Workout) (*domain.Workout, error)

	DeleteWorkout(ctx context.Context, id string) error

	ListWorkouts(ctx context.Context, user *domain.User, query domain.GetWorkoutListQuery) ([]domain.Workout, int64, error)
}

type AchievementRepository interface {
	CreateAchievement(ctx context.Context, achievement domain.Achievement) (domain.Achievement, error)

	GetAchievementById(ctx context.Context, id string) (*domain.Achievement, error)

	UpdateAchievement(ctx context.Context, achievement domain.Achievement) (*domain.Achievement, error)

	DeleteAchievement(ctx context.Context, id string) error

	ListAchievements(ctx context.Context, user *domain.User, pagination domain.PaginationQuery) ([]domain.Achievement, int64, error)
}

type GoalRepository interface {
	CreateGoal(ctx context.Context, goal domain.Goal) (*domain.Goal, error)

	GetGoalById(ctx context.Context, id string) (*domain.Goal, error)

	UpdateGoal(ctx context.Context, goal domain.Goal) (*domain.Goal, error)

	DeleteGoal(ctx context.Context, id string) error

	ListGoalsByUserId(ctx context.Context, user *domain.User) ([]domain.Goal, error)

	ListGoals(ctx context.Context, user *domain.User, query domain.GetGoalListQuery) ([]domain.Goal, int64, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)

	GetUserByID(ctx context.Context, id string) (*domain.User, error)

	GetUser(ctx context.Context, username, password string) (*domain.User, error)

	UpdateUser(ctx context.Context, user domain.User) (*domain.User, error)

	DeleteUser(ctx context.Context, id string) error

	ListUsers(ctx context.Context) ([]domain.User, error)
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

func filterDateQueries(dateFrom, dateTo time.Time, fieldName string, filter bson.M) error {
	if !dateFrom.IsZero() && !dateTo.IsZero() {
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			"$and": []bson.M{
				{fieldName: bson.M{"$gte": dateFrom}},
				{fieldName: bson.M{"$lte": dateTo}},
			},
		})
	}

	if !dateFrom.IsZero() && dateTo.IsZero() {
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			fieldName: bson.M{"$gte": dateFrom},
		})
	}

	if dateFrom.IsZero() && !dateTo.IsZero() {
		filter["$and"] = append(filter["$and"].([]bson.M), bson.M{
			fieldName: bson.M{"$lte": dateTo},
		})
	}

	return nil
}
