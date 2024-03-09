package services

import (
	"context"
	"log"

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

func (s *WorkoutService) GetUserWorkout(ctx context.Context, userId string) []models.Workout {
	workouts, err := s.rep.WorkoutRepository.GetWorkoutByUserId(
		ctx,
		userId,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return workouts
}

func (s *WorkoutService) CreateWorkout(ctx context.Context, userId string) []models.Workout {
	workouts, err := s.rep.WorkoutRepository.GetWorkoutByUserId(
		ctx,
		ctx.Value("user").(string),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return workouts
}
