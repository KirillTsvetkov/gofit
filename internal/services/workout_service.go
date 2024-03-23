package services

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
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

func (s *WorkoutService) GetUserWorkout(ctx context.Context, user *domain.User, query domain.GetWorkoutListQuery) ([]domain.Workout, int64) {
	workouts, total, err := s.rep.WorkoutRepository.ListWorkouts(
		ctx,
		user,
		query.PaginationQuery,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	if workouts == nil {
		workouts = []domain.Workout{}
	}

	return workouts, total
}

func (s *WorkoutService) CreateWorkout(ctx context.Context, exercise []domain.Exercise, user *domain.User) *domain.Workout {

	workout := &domain.Workout{
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
