package orders

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/oganes5796/shops-list/internal/orders/model"
	"github.com/oganes5796/shops-list/internal/orders/repository"
	"github.com/oganes5796/shops-list/internal/orders/repository/orders"
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

func (s *ordersService) Create(ctx context.Context, req *model.CreateOrderRequest) (uuid.UUID, error) {

	var (
		totalPrice  float64
		totalVolume int
	)

	for _, item := range req.Cart {
		totalPrice += float64(item.Quantity) * item.Price
		totalVolume += item.Quantity
	}

	order := &model.OrderInfo{
		ID:            uuid.New(),
		UserID:        req.UserID,
		ShopID:        req.ShopID,
		Cart:          req.Cart,
		SummaryVolume: totalVolume,
		Price:         totalPrice,
		Status:        model.OrderStatusCreated,
	}

	return s.repo.Create(ctx, order)
}

func (s *ordersService) GetByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, orders.ErrOrderNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}

	return order, nil
}

func (s *ordersService) Update(ctx context.Context, id uuid.UUID, req *model.UpdateOrderRequest) error {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, orders.ErrOrderNotFound) {
			return ErrOrderNotFound
		}
		return err
	}

	if req.Status != nil {
		if !isValidStatusTransition(order.Info.Status, *req.Status) {
			return ErrInvalidStatusTransition
		}
		order.Info.Status = *req.Status
	}

	return s.repo.Update(ctx, &order.Info)
}

func isValidStatusTransition(from, to model.OrderStatus) bool {
	switch from {
	case model.OrderStatusCreated:
		return to == model.OrderStatusInProgress
	case model.OrderStatusInProgress:
		return to == model.OrderStatusSuccess || to == model.OrderStatusFail
	default:
		return false
	}
}
