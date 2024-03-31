package services

import (
	"context"
	"log"
	"time"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		query,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	if workouts == nil {
		workouts = []domain.Workout{}
	}

	return workouts, total
}

func (s *WorkoutService) CreateWorkout(ctx context.Context, date time.Time, exercise []domain.Exercise, user *domain.User) *domain.Workout {
	workout := &domain.Workout{
		UserID:    user.ID,
		Date:      date,
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

func (s *WorkoutService) UpdateWorkout(ctx context.Context, user *domain.User, id primitive.ObjectID, query domain.UpdateWorkoutQuery) (domain.Workout, error) {
	var workout domain.Workout
	workout, err := s.rep.WorkoutRepository.GetWorkoutById(ctx, user, id)
	if err != nil {
		return workout, err
	}

	updatedWorkout, err := s.rep.WorkoutRepository.UpdateWorkout(
		ctx,
		workout,
		query,
	)

	if err != nil {
		return workout, err
	}

	return updatedWorkout, err
}
