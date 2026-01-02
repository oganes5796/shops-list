package service

import (
	"context"

	"github.com/oganes5796/shops-list/internal/shops-list/model"
)

type AuthService interface {
	Register(ctx context.Context, user *model.UserInfo) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
}

type ShopService interface {
	Create(ctx context.Context, info *model.ShopInfo) (int64, error)
	GetAll(ctx context.Context) ([]model.Shop, error)
	GetByID(ctx context.Context, idShop int64) (*model.Shop, error)
	Update(ctx context.Context, idShop int64, info *model.ShopInfo) error
	Delete(ctx context.Context, idShop int64) error
}
