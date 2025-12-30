package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
)

const userTable = "users"

type authRepository struct {
	conn *pgx.Conn
}

func NewAuthRepository(conn *pgx.Conn) repository.AuthRepository {
	return &authRepository{
		conn: conn,
	}
}

func (r *authRepository) CreateUser(ctx context.Context, user *model.UserInfo) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (username, role) VALUES ($1, $2) RETURNING id", userTable)
	err := r.conn.QueryRow(ctx, query, user.Username, user.Role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository:auth:Create:row.Scan: %w", err)
	}
	return id, nil
}

func (r *authRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user := &model.User{
		Info: model.UserInfo{},
	}
	query := fmt.Sprintf("SELECT id, username, role, created_at FROM %s WHERE username = $1", userTable)
	row := r.conn.QueryRow(ctx, query, username)
	err := row.Scan(&user.ID, &user.Info.Username, &user.Info.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUsernameNotFound
		}

		return nil, fmt.Errorf("repository:auth:GetUserByUsername:row.Scan: %w", err)
	}
	return user, nil
}
