package serv

import (
	"github.com/oganes5796/shops-list/internal/shops-list/repository/repo"
	"github.com/oganes5796/shops-list/internal/shops-list/service"
	"github.com/oganes5796/shops-list/internal/shops-list/service/auth"
	"github.com/oganes5796/shops-list/internal/shops-list/service/shop"
)

type Serv struct {
	service.AuthService
	service.ShopService
}

func NewService(repos *repo.Repository) *Serv {
	return &Serv{
		AuthService: auth.NewAuthService(repos.AuthRepository),
		ShopService: shop.NewShopService(repos.ShopRepository),
	}
}
