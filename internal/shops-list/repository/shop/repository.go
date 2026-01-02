package shop

import (
	"github.com/jackc/pgx/v5"
	"github.com/oganes5796/shops-list/internal/shops-list/repository"
)

const (
	tableName = "shops"
)

type shopRepository struct {
	conn *pgx.Conn
}

func NewShopRepository(conn *pgx.Conn) repository.ShopRepository {
	return &shopRepository{
		conn: conn,
	}
}
