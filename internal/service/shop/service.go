package shop

import (
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
