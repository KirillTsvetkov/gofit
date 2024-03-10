package services

import (
	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/pkg/auth"
)

type Service struct {
	*AuthService
	*WorkoutService
}

func NewRepository(rep *repository.Repository, jwtManager *auth.Manager) *Service {
	return &Service{
		AuthService:    NewAuthService(rep, jwtManager),
		WorkoutService: NewWorkoutService(rep),
	}
}
