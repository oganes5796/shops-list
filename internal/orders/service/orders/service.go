package orders

import (
	"context"

	"github.com/oganes5796/shops-list/internal/orders/model"
	"github.com/oganes5796/shops-list/internal/orders/repository"
	"github.com/oganes5796/shops-list/internal/orders/service"
)

type ordersService struct {
	repo repository.OrderRepository
}

func NewOrdersService(repo repository.OrderRepository) service.OrdersService {
	return &ordersService{
		repo: repo,
	}
}

func (s *ordersService) Create(ctx context.Context, order *model.OrderInfo) (string, error) {
	return s.repo.Create(ctx, order)
}

func (s *ordersService) GetByID(ctx context.Context, idOrder string) (*model.Order, error) {
	return s.repo.GetByID(ctx, idOrder)
}

func (s *ordersService) Update(ctx context.Context, idOrder string, order *model.OrderInfo) error {
	return s.repo.Update(ctx, idOrder, order)
}
