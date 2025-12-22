package repository

import (
	"context"

	"github.com/oganes5796/shops-list/internal/model"
)

type ShopRepository interface {
	Create(ctx context.Context, info *model.ShopInfo) (int64, error)
	GetAll(ctx context.Context) ([]model.Shop, error)
	GetByID(ctx context.Context, id int64) (*model.Shop, error)
	Update(ctx context.Context, id int64, info *model.ShopInfo) error
	Delete(ctx context.Context, idShop int64) error
}
