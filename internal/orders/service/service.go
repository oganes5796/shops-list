package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/oganes5796/shops-list/internal/orders/model"
)

type OrdersService interface {
	Create(ctx context.Context, req *model.CreateOrderRequest) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	Update(ctx context.Context, id uuid.UUID, req *model.UpdateOrderRequest) error
}
