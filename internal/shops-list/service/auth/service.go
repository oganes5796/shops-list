package auth

import (
	"context"
	"errors"

	"github.com/oganes5796/shops-list/internal/shops-list/model"
	"github.com/oganes5796/shops-list/internal/shops-list/repository"
	"github.com/oganes5796/shops-list/internal/shops-list/repository/auth"
	"github.com/oganes5796/shops-list/internal/shops-list/service"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Register(ctx context.Context, user *model.UserInfo) (int64, error) {
	if user.Username == "" {
		return 0, ErrEmptyUsername
	}
	if user.Role != model.RoleUser && user.Role != model.RoleManager {
		return 0, ErrInvalidRole
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *authService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, auth.ErrUsernameNotFound) {
			return nil, ErrUsernameNotFound
		}
		return nil, err
	}

	return user, nil
}
