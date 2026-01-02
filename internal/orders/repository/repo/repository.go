package repo

import (
	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/orders/repository"
	"github.com/oganes5796/shops-list/internal/orders/repository/orders"
)

type Repository struct {
	repository.OrderRepository
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		OrderRepository: orders.NewOrdersRepository(conn),
	}
}
