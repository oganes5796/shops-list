package repo

import (
	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/repository"
	"github.com/oganes5796/shops-list/internal/repository/auth"
	"github.com/oganes5796/shops-list/internal/repository/shop"
)

type Repository struct {
	repository.AuthRepository
	repository.ShopRepository
}

func NewRepository(conn *pgx.Conn) *Repository {
	return &Repository{
		AuthRepository: auth.NewAuthRepository(conn),
		ShopRepository: shop.NewShopRepository(conn),
	}
}
