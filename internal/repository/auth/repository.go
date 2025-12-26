package auth

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
)

type authRepository struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) repository.AuthRepository {
	return &authRepository{
		conn: conn,
	}
}

func (r *authRepository) CreateUser(ctx context.Context, user *model.UserInfo) (int64, error) {
	// Implementation here
	return 0, nil
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	// Implementation here
	return nil, nil
}
