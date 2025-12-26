package auth

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
	"github.com/oganes5796/shops-list/internal/service"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) CreateUser(ctx context.Context, user *model.UserInfo) (int64, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *authService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}
