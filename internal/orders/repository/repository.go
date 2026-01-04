package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/oganes5796/shops-list/internal/orders/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.OrderInfo) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	Update(ctx context.Context, order *model.OrderInfo) error
}
