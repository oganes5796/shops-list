package shop

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
)

const (
	tableName = "shops"
)

type shopRepository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) repository.ShopRepository {
	return &shopRepository{
		conn: conn,
	}
}

func (r *shopRepository) GetAll(ctx context.Context) ([]model.Shop, error) {
	// Implementation here
	return nil, nil
}

func (r *shopRepository) Update(ctx context.Context, id int64, info *model.ShopInfo) error {
	// Implementation here
	return nil
}
