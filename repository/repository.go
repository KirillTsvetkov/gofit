package repository

import "github.com/KirillTsvetkov/gofit/models"

type WorkoutRepository interface {
	CreateWorkout(workout models.Workout) (string, error)

	GetWorkoutById(id string) (models.Workout, error)

	UpdateWorkout(workout models.Workout) error

	DeleteWorkout(id string) error

	ListWorkoutsByUserId(userId int) ([]models.Workout, error)
}

type AchievementRepository interface {
	CreateAchievement(achievement models.Achievement) (string, error)

	GetAchievementById(id string) (models.Achievement, error)

	UpdateAchievement(achievement models.Achievement) error

	DeleteAchievement(id string) error

	ListAchievementsByUserId(userId int) ([]models.Achievement, error)

	ListAchievementsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Achievement, error)
}

type GoalRepository interface {
	CreateGoal(goal models.Goal) (string, error)

	GetGoalById(id string) (models.Goal, error)

	UpdateGoal(goal models.Goal) error

	DeleteGoal(id string) error

	ListGoalsByUserId(userId int) ([]models.Goal, error)

	ListGoalsByUserIdAndWorkoutId(userId, workoutId int) ([]models.Goal, error)
}

type Repository struct {
}
