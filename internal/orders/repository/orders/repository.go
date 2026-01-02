package orders

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/orders/model"
	"github.com/oganes5796/shops-list/internal/orders/repository"
)

const (
	tableName = "orders"
)

type ordersRepository struct {
	conn *pgx.Conn
}

func NewOrdersRepository(conn *pgx.Conn) repository.OrderRepository {
	return &ordersRepository{
		conn: conn,
	}
}

func (r *ordersRepository) Create(ctx context.Context, order *model.OrderInfo) (string, error) {
	// Implementation goes here
	return "", nil
}

func (r *ordersRepository) GetByID(ctx context.Context, idOrder string) (*model.Order, error) {
	// Implementation goes here
	return nil, nil
}

func (r *ordersRepository) Update(ctx context.Context, idOrder string, order *model.OrderInfo) error {
	// Implementation goes here
	return nil
}
