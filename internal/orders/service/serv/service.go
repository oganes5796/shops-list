package serv

import (
	"github.com/oganes5796/shops-list/internal/orders/repository/repo"
	"github.com/oganes5796/shops-list/internal/orders/service"
	"github.com/oganes5796/shops-list/internal/orders/service/orders"
)

type Serv struct {
	service.OrdersService
}

func NewService(repos *repo.Repository) *Serv {
	return &Serv{
		OrdersService: orders.NewOrdersService(repos.OrderRepository),
	}
}
