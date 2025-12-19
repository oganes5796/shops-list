package shop

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
)

type shopRepository struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) repository.ShopRepository {
	return &shopRepository{
		conn: conn,
	}
}

func (r *shopRepository) Create(ctx context.Context, info *model.ShopInfo) (int64, error) {
	// Implementation here
	return 0, nil
}

func (r *shopRepository) GetAll(ctx context.Context) ([]model.Shop, error) {
	// Implementation here
	return nil, nil
}

func (r *shopRepository) GetByID(ctx context.Context, id int64) (*model.Shop, error) {
	// Implementation here
	return nil, nil
}

func (r *shopRepository) Update(ctx context.Context, id int64, info *model.ShopInfo) error {
	// Implementation here
	return nil
}

func (r *shopRepository) Delete(ctx context.Context, id int64) error {
	// Implementation here
	return nil
}
