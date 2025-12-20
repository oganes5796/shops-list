package shop

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
	"github.com/oganes5796/shops-list/internal/repository"
	"github.com/oganes5796/shops-list/internal/service"
)

type shopService struct {
	repo repository.ShopRepository
}

func NewService(repo repository.ShopRepository) service.ShopService {
	return &shopService{
		repo: repo,
	}
}

func (r *shopService) Create(ctx context.Context, info *model.ShopInfo) (int64, error) {
	// Implementation here
	return 0, nil
}

func (r *shopService) GetAll(ctx context.Context) ([]model.Shop, error) {
	// Implementation here
	return nil, nil
}

func (r *shopService) GetByID(ctx context.Context, id int64) (*model.Shop, error) {
	// Implementation here
	return nil, nil
}

func (r *shopService) Update(ctx context.Context, id int64, info *model.ShopInfo) error {
	// Implementation here
	return nil
}

func (r *shopService) Delete(ctx context.Context, id int64) error {
	// Implementation here
	return nil
}
