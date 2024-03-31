package services

import (
	"context"
	"log"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GoalService struct {
	rep *repository.Repository
}

func NewGoalService(rep *repository.Repository) *GoalService {
	return &GoalService{
		rep: rep,
	}
}

func (s *GoalService) GetUserGoals(ctx context.Context, user *domain.User, query domain.GetGoalListQuery) ([]domain.Goal, int64) {
	goals, total, err := s.rep.GoalRepository.ListGoals(
		ctx,
		user,
		query,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	if goals == nil {
		goals = []domain.Goal{}
	}

	return goals, total
}

func (s *GoalService) GetGoal(ctx context.Context, user *domain.User, id primitive.ObjectID) (domain.Goal, error) {
	goal, err := s.rep.GoalRepository.GetGoalById(ctx, user, id)
	log.Print(goal)
	if err != nil {
		return goal, err
	}

	return goal, nil
}

func (s *GoalService) UpdateGoal(ctx context.Context, user *domain.User, id primitive.ObjectID, query domain.UpdateGoalQuery) (domain.Goal, error) {
	var goal domain.Goal
	goal, err := s.rep.GoalRepository.GetGoalById(ctx, user, id)
	if err != nil {
		return goal, err
	}

	updatedGoal, err := s.rep.GoalRepository.UpdateGoal(
		ctx,
		goal,
		query,
	)

	if err != nil {
		return goal, err
	}

	return updatedGoal, err
}
