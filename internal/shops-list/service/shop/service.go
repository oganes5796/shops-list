package shop

import (
	"github.com/oganes5796/shops-list/internal/shops-list/repository"
	"github.com/oganes5796/shops-list/internal/shops-list/service"
)

type shopService struct {
	repo repository.ShopRepository
}

func NewShopService(repo repository.ShopRepository) service.ShopService {
	return &shopService{
		repo: repo,
	}
}
