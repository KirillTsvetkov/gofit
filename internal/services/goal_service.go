package services

import (
	"context"
	"log"

	"github.com/KirillTsvetkov/gofit/internal/domain"
	"github.com/KirillTsvetkov/gofit/internal/repository"
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
