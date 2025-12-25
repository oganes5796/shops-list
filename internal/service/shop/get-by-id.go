package shop

import (
	"context"
	"errors"

	"github.com/oganes5796/shops-list/internal/model"
	repoShop "github.com/oganes5796/shops-list/internal/repository/shop"
)

func (r *shopService) GetByID(ctx context.Context, idShop int64) (*model.Shop, error) {
	shop, err := r.repo.GetByID(ctx, idShop)
	if err != nil {
		if errors.Is(err, repoShop.ErrShopNotFound) {
			return nil, ErrShopNotFound
		}
		return nil, err
	}

	return shop, nil
}
