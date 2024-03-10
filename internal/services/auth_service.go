package services

import (
	"context"
	"log"

	"github.com/KirillTsvetkov/gofit/internal/repository"
	"github.com/KirillTsvetkov/gofit/pkg/auth"
)

type AuthService struct {
	jwtManager *auth.Manager
	rep        *repository.Repository
}

func NewAuthService(rep *repository.Repository, jwtManager *auth.Manager) *AuthService {
	return &AuthService{
		jwtManager: jwtManager,
		rep:        rep,
	}
}

func (s *AuthService) CreateToken(ctx context.Context, email string, password string) (string, error) {
	user, err := s.rep.UserRepository.GetUser(ctx, email, password)
	if err != nil {
		log.Fatal(err.Error())
	}
	stringObjectID := user.ID.Hex()
	token, err := s.jwtManager.GenerateJWT(stringObjectID, 300)

	return token, err
}
