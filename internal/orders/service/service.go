package service

import (
	"context"

	"github.com/oganes5796/shops-list/internal/orders/model"
)

type OrdersService interface {
	Create(ctx context.Context, order *model.OrderInfo) (string, error)
	GetByID(ctx context.Context, idOrder string) (*model.Order, error)
	Update(ctx context.Context, idOrder string, order *model.OrderInfo) error
}
