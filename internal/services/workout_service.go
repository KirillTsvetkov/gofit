package services

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/models"
	"github.com/KirillTsvetkov/gofit/internal/repository"
)

type WorkoutService struct {
	rep *repository.Repository
}

func NewWorkoutService(rep *repository.Repository) *WorkoutService {
	return &WorkoutService{
		rep: rep,
	}
}

func (s *WorkoutService) GetUserWorkout(ctx context.Context, userId *models.User) []models.Workout {
	workouts, err := s.rep.WorkoutRepository.ListWorkouts(
		ctx,
		userId,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	if workouts == nil {
		workouts = []models.Workout{}
	}

	return workouts
}

func (s *WorkoutService) CreateWorkout(ctx context.Context, exercise []models.Exercise, user *models.User) *models.Workout {
	workout := &models.Workout{
		UserID:    user.ID,
		Date:      time.Now(),
		Exercises: exercise,
	}

	result, err := s.rep.WorkoutRepository.CreateWorkout(
		ctx,
		workout,
		user,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
