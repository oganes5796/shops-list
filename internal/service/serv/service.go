package serv

import (
	"github.com/oganes5796/shops-list/internal/repository/repo"
	"github.com/oganes5796/shops-list/internal/service"
	"github.com/oganes5796/shops-list/internal/service/auth"
	"github.com/oganes5796/shops-list/internal/service/shop"
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
